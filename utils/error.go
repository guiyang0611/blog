package utils

import (
	"blog/log"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"runtime/debug"
)

var (
	ErrUnauthorized = errors.New("未授权")
	ErrNotFound     = errors.New("资源不存在")
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 打印错误和堆栈信息
				stack := debug.Stack()
				log.Logger.Error("Panic", zap.Any("error", err), zap.String("stack", string(stack)))

				Error(c, http.StatusInternalServerError, "服务器内部错误")
			}
		}()

		c.Next()

		// 处理非 panic 错误（如业务错误）
		for _, err := range c.Errors {
			log.Logger.Error("请求错误", zap.Error(err.Err))
			switch {
			case errors.Is(err.Err, gorm.ErrRecordNotFound):
				Error(c, http.StatusNotFound, "资源不存在")
			case errors.Is(err.Err, ErrUnauthorized):
				Error(c, http.StatusUnauthorized, "未授权")
			default:
				Error(c, http.StatusInternalServerError, "系统错误")
			}
		}
	}
}
