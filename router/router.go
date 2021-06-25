package router

import (
	"github.com/gin-gonic/gin"
	v1 "mofan-blog/api/v1"
	"mofan-blog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router_v1 := r.Group("api/v1")
	{
		// User 模块的路由接口
		router_v1.POST("user/add", v1.AddUser)
		router_v1.GET("users", v1.GetUsers)
		router_v1.PUT("user/:id", v1.EditUser)
		router_v1.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		// 文章模块的路由接口
		router_v1.GET("hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "ok",
			})

		})
	}
	r.Run(utils.HttpPort)
}
