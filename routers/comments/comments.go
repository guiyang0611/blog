package comments

import (
	"blog/config"
	"blog/log"
	"blog/models"
	"blog/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func Save(c *gin.Context) {
	var commentDto models.Comment
	currUser, _ := c.Get("user")
	log.Logger.Info("当前登录用户信息", zap.Any("currUser:", currUser))
	if currUser == nil {
		utils.Error(c, http.StatusUnauthorized, "用户未登录")
	}
	if err := c.ShouldBindJSON(&commentDto); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	commentDto.UserID = currUser.(models.User).ID
	config.GormDB.Create(&commentDto)
	utils.SUCCESS(c, "添加成功!", nil)
}

// Query 获取文章的所有评论,无需登录
func Query(c *gin.Context) {
	var commentDto models.Comment
	if err := c.ShouldBindJSON(&commentDto); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	log.Logger.Info("查询参数为", zap.Any("commentDto:", commentDto))
	if commentDto.PostID == 0 {
		utils.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var comments []models.Comment
	config.GormDB.Debug().Where(&commentDto).Find(&comments)
	utils.SUCCESS(c, "查询成功!", comments)
}

func Update(c *gin.Context) {
	var commentDto models.Comment
	currUser, _ := c.Get("user")
	log.Logger.Info("当前登录用户信息", zap.Any("currUser:", currUser))
	if currUser == nil {
		utils.Error(c, http.StatusUnauthorized, "用户未登录")
	}
	loginUser := currUser.(*models.User)
	if err := c.ShouldBindJSON(&commentDto); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	if commentDto.ID == 0 {
		fmt.Println("id不能为空!")
		utils.Error(c, http.StatusBadRequest, "id不能为空")
		return
	}
	var user models.User
	config.GormDB.First(&user, commentDto.UserID)
	if user.ID != loginUser.ID {
		utils.Error(c, http.StatusUnauthorized, "暂无权限修改!")
	}
	config.GormDB.Save(&commentDto)
	utils.SUCCESS(c, "更新成功!", nil)
}

func Delete(c *gin.Context) {
	var comment models.Comment
	currUser, _ := c.Get("user")
	log.Logger.Info("当前登录用户信息", zap.Any("currUser:", currUser))
	if currUser == nil {
		utils.Error(c, http.StatusUnauthorized, "用户未登录")
	}
	loginUser := currUser.(*models.User)
	if err := c.ShouldBindJSON(&comment); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var user models.User
	config.GormDB.First(&user, comment.UserID)
	if user.ID != loginUser.ID {
		utils.Error(c, http.StatusUnauthorized, "暂无权限删除!")
	}
	config.GormDB.Delete(&models.Post{}, comment.ID)
	utils.SUCCESS(c, "删除成功!", nil)
}
