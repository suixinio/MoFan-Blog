package v1

import (
	"github.com/gin-gonic/gin"
	"mofan-blog/middleware"
	"mofan-blog/model"
	"mofan-blog/utils/errmsg"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	var token string
	var code int

	c.ShouldBindJSON(&data)
	code = model.CheckLogin(data.UserName, data.Password)
	if code == errmsg.SUCCESS {
		token, _ = middleware.SetToken(data.UserName)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})

}
