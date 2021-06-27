package router

import (
	"github.com/gin-gonic/gin"
	"mofan-blog/api/v1"
	"mofan-blog/middleware"
	"mofan-blog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	authV1 := r.Group("api/v1")
	authV1.Use(middleware.JwtToken())
	{
		// User 模块的路由接口
		authV1.PUT("user/:id", v1.EditUser)
		authV1.DELETE("user/:id", v1.DeleteUser)

		// 分类模块的路由接口
		authV1.POST("category/add", v1.AddCategory)
		authV1.PUT("category/:id", v1.EditCate)
		authV1.DELETE("category/:id", v1.DeleteCate)

		// 文章模块的路由接口
		authV1.POST("article/add", v1.AddArticle)
		authV1.PUT("article/:id", v1.EditArt)
		authV1.DELETE("article/:id", v1.DeleteArt)

		// 上传文件
		authV1.POST("upload/", v1.Upload)
	}
	routerV1 := r.Group("api/v1")
	{
		// 登陆
		routerV1.POST("login", v1.Login)

		// User 模块的路由接口
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("users", v1.GetUsers)

		// 分类模块的路由接口
		routerV1.GET("category", v1.GetCate)
		// 文章模块的路由接口
		routerV1.GET("article", v1.GetArt)
		routerV1.GET("article/list/:id", v1.GetCateArt)
		routerV1.GET("article/:id", v1.GetArtInfo)
	}

	r.Run(utils.HttpPort)
}
