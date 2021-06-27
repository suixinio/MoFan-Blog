package v1

import (
	"github.com/gin-gonic/gin"
	"mofan-blog/model"
	"mofan-blog/utils/errmsg"
	"net/http"
)

func Upload(c *gin.Context) {

	file, fileHander, _ := c.Request.FormFile("file")

	fileSize := fileHander.Size

	url, code := model.UploadFile(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
