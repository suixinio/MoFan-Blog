package router

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"mofan-blog/api/v1"
	"mofan-blog/middleware"
	"mofan-blog/utils"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("admin", "web/admin/dist/index.html")
	r.AddFromFiles("front", "web/front/dist/index.html")

	return r
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()

	r.HTMLRender = createMyRender()

	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	//r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

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
		authV1.POST("upload", v1.Upload)

		// 更新个人设置
		authV1.PUT("profile/:id", v1.UpdateProfile)
	}
	routerV1 := r.Group("api/v1")
	{
		// 登陆
		routerV1.POST("login", v1.Login)

		// User 模块的路由接口
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("users", v1.GetUsers)
		routerV1.GET("user/:id", v1.GetUserInfo)

		// 分类模块的路由接口
		routerV1.GET("category/:id", v1.GetCateInfo)
		routerV1.GET("category", v1.GetCate)

		// 文章模块的路由接口
		routerV1.GET("article", v1.GetArt)
		routerV1.GET("article/list/:id", v1.GetCateArt)
		routerV1.GET("article/:id", v1.GetArtInfo)

		// 获取用户用户信息
		routerV1.GET("profile/:id", v1.GetProfile)
	}

	r.Run(utils.HttpPort)
}
