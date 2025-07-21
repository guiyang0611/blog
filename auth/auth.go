package auth

import (
	"blog/config"
	"blog/log"
	"blog/models"
	"blog/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func AuthInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Logger.Warn("缺少Token")
			utils.Error(c, http.StatusUnauthorized, "未授权")
			c.Abort()
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Logger.Warn("Token格式错误")
			utils.Error(c, http.StatusUnauthorized, "无效Token")
			c.Abort()
			return
		}
		tokenString := parts[1]
		// 解析 token
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			jwtSecret := []byte(config.Cfg.SecretKey)
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			log.Logger.Warn("Token无效", zap.Error(err))
			utils.Error(c, http.StatusUnauthorized, "无效Token")
			c.Abort()
			return
		}
		// 查询用户
		var user models.User
		if err := config.GormDB.Where("username = ?", claims.Username).First(&user).Error; err != nil {
			log.Logger.Warn("用户不存在", zap.Any("username", claims.Username))
			utils.Error(c, http.StatusUnauthorized, "用户不存在")
			c.Abort()
			return
		}
		// 把用户信息放入 context
		c.Set("user", &user)
		c.Next()
	}
}
