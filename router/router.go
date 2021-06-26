package router

import (
	"github.com/gin-gonic/gin"
	"mofan-blog/api/v1"
	"mofan-blog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	routerV1 := r.Group("api/v1")
	{
		// User 模块的路由接口
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("users", v1.GetUsers)
		routerV1.PUT("user/:id", v1.EditUser)
		routerV1.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		// 文章模块的路由接口
		routerV1.GET("hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "ok",
			})

		})
	}

	r.Run(utils.HttpPort)
}
