package posts

import (
	"blog/config"
	"blog/log"
	"blog/models"
	"blog/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func Save(c *gin.Context) {
	var post models.Post
	currUser, _ := c.Get("user")
	log.Logger.Info("当前登录用户信息", zap.Any("currUser:", currUser))
	if currUser == nil {
		utils.Error(c, http.StatusUnauthorized, "用户未登录")
	}
	if err := c.ShouldBindJSON(&post); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	post.UserID = currUser.(models.User).ID
	config.GormDB.Create(&post)
	utils.SUCCESS(c, "添加成功!", nil)
}

func Query(c *gin.Context) {
	var postDto models.Post
	if err := c.ShouldBindJSON(&postDto); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	postJson, _ := json.Marshal(postDto)
	fmt.Println("查询参数为:", string(postJson))
	var posts []models.Post
	config.GormDB.Debug().Where(&postDto).Find(&posts)
	utils.SUCCESS(c, "查询成功!", posts)

}

func Update(c *gin.Context) {
	var postInput models.Post
	currUser, exists := c.Get("user")
	log.Logger.Info("当前登录用户信息", zap.Any("currUser:", currUser))
	if !exists || currUser == nil {
		utils.Error(c, http.StatusUnauthorized, "用户未登录")
	}
	loginUser, _ := currUser.(*models.User)
	if err := c.ShouldBindJSON(&postInput); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	if postInput.ID == 0 {
		fmt.Println("id不能为空!")
		utils.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var post models.Post
	config.GormDB.First(&post, postInput.ID)
	var user models.User
	if err := config.GormDB.First(&user, post.UserID).Error; err != nil {
		log.Logger.Error("该条文章所属用户不存在", zap.Error(err))
		utils.Error(c, http.StatusUnauthorized, "系统异常")
		return
	}
	if user.ID != loginUser.ID {
		utils.Error(c, http.StatusUnauthorized, "暂无权限修改!")
		return
	}
	config.GormDB.Model(&postInput).Updates(&postInput)
	utils.SUCCESS(c, "更新成功!", nil)
}

func Delete(c *gin.Context) {
	var post models.Post
	currUser, _ := c.Get("user")
	log.Logger.Info("当前登录用户信息", zap.Any("currUser:", currUser))
	if currUser == nil {
		utils.Error(c, http.StatusUnauthorized, "用户未登录")
	}
	loginUser := currUser.(*models.User)
	if err := c.ShouldBindJSON(&post); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var user models.User
	config.GormDB.First(&user, post.UserID)
	if user.ID != loginUser.ID {
		utils.Error(c, http.StatusUnauthorized, "暂无权限删除!")
	}
	config.GormDB.Delete(&models.Post{}, post.ID)
	utils.SUCCESS(c, "删除成功!", nil)
}
