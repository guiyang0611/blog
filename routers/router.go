package routers

import (
	"blog/auth"
	"blog/routers/comments"
	"blog/routers/posts"
	"blog/routers/user"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		// 不需要登录的接口
		noAuth := api.Group("")
		userGroup := noAuth.Group("/user")
		noAuthPosts := noAuth.Group("/posts")
		noAuthComment := noAuth.Group("/comment")
		{
			userGroup.POST("/register", user.Register)
			userGroup.POST("/login", user.Login)
			noAuthPosts.POST("/query", posts.Query)
			noAuthComment.POST("/query", posts.Query)
		}

		// 业务接口，需要登录
		authorized := api.Group("")
		authorized.Use(auth.AuthInterceptor())
		postsGroup := authorized.Group("/posts")
		{
			postsGroup.POST("/save", posts.Save)
			postsGroup.POST("/update", posts.Update)
			postsGroup.POST("/delete", posts.Delete)
		}
		commentGroup := authorized.Group("/comment")
		{
			commentGroup.POST("/save", comments.Save)
			commentGroup.POST("/update", comments.Update)
			commentGroup.POST("/delete", comments.Delete)
		}
	}
	return r
}
