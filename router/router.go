package router

import (
	"github.com/gin-gonic/gin"
	"mofan-blog/api/v1"
	"mofan-blog/middleware"
	"mofan-blog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	authvV1 := r.Group("api/v1")
	authvV1.Use(middleware.JwtToken())
	{
		// User 模块的路由接口
		authvV1.PUT("user/:id", v1.EditUser)
		authvV1.DELETE("user/:id", v1.DeleteUser)

		// 分类模块的路由接口
		authvV1.POST("category/add", v1.AddCategory)
		authvV1.PUT("category/:id", v1.EditCate)
		authvV1.DELETE("category/:id", v1.DeleteCate)

		// 文章模块的路由接口
		authvV1.POST("article/add", v1.AddArticle)
		authvV1.PUT("article/:id", v1.EditArt)
		authvV1.DELETE("article/:id", v1.DeleteArt)
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
