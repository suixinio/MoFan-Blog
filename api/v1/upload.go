package v1

import (
	"github.com/gin-gonic/gin"
	"mofan-blog/service/upload"
	//"mofan-blog/service/upload/qiniu"
	"mofan-blog/utils/errmsg"
	"net/http"
)

func Upload(c *gin.Context) {

	file, fileHandler, _ := c.Request.FormFile("file")
	if fileHandler == nil {
		code := errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
	url, code := upload.Upup.UploadFile(file, fileHandler)
	//url, code := qiniu.UploadFile(file, fileHandler)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
