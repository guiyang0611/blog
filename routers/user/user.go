package user

import (
	"blog/auth"
	"blog/config"
	"blog/log"
	"blog/models"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context) {
	var userInput models.User
	db := config.GormDB
	// 1. 绑定参数
	if err := c.ShouldBindJSON(&userInput); err != nil {
		log.Logger.Warn("注册失败：参数绑定错误", zap.Error(err))
		utils.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	log.Logger.Info("注册参数为", zap.Any("userInput:", userInput))
	if len(userInput.Username) == 0 {
		utils.Error(c, http.StatusBadRequest, "用户名不能为空")
	}
	// 2. 密码长度检查
	if len(userInput.Password) < 6 {
		log.Logger.Warn("注册失败：密码过短", zap.String("username", userInput.Username))
		utils.Error(c, http.StatusBadRequest, "密码至少需要6位")
		return
	}
	// 3. 检查用户名是否已存在
	var existingUser models.User
	if err := db.Where("username = ?", userInput.Username).First(&existingUser).Error; err == nil {
		log.Logger.Warn("注册失败：用户名已存在", zap.String("username", userInput.Username))
		utils.Error(c, http.StatusConflict, "用户名已存在")
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Logger.Error("注册失败：密码加密失败", zap.Error(err))
		utils.Error(c, http.StatusInternalServerError, "账号或密码错误")
		return
	}
	userInput.Password = string(hashedPassword)
	if err := db.Create(&userInput).Error; err != nil {
		log.Logger.Error("注册失败：数据库异常", zap.Error(err))
		utils.Error(c, http.StatusInternalServerError, "账号或密码错误")
		return
	}
	log.Logger.Info("注册成功", zap.String("username", userInput.Username))
	utils.SUCCESS(c, "注册成功", nil)
}

func Login(c *gin.Context) {
	var userInput models.User
	db := config.GormDB
	if err := c.ShouldBindJSON(&userInput); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())

		return
	}
	var storedUser models.User
	if err := db.Where("username = ?", userInput.Username).First(&storedUser).Error; err != nil {
		log.Logger.Warn("登录失败：用户不存在", zap.String("username", userInput.Username))
		utils.Error(c, http.StatusUnauthorized, "用户名或密码错误")
		return
	}
	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(userInput.Password)); err != nil {
		log.Logger.Warn("登录失败：密码错误", zap.String("username", userInput.Username))
		utils.Error(c, http.StatusUnauthorized, "用户名或密码错误")
		return
	}
	// 生成 JWT
	token, err := auth.GenerateToken(userInput.Username, userInput.Password)
	if err != nil {
		log.Logger.Error("登录失败：生成Token失败", zap.Error(err))
		utils.Error(c, http.StatusUnauthorized, "登录失败")
		return
	}
	log.Logger.Info("登录成功", zap.String("token", token))
	utils.SUCCESS(c, "登录成功!", token)

}
