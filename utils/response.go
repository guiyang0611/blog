package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	OK             = 0
	ERROR          = 1
	INVALID_PARAMS = 400
)

func SUCCESS(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, Response{
		Code:    1,
		Message: message,
	})
}
