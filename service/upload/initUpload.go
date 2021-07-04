package upload

import (
	"mime/multipart"
	"mofan-blog/service/upload/gsky"
	"mofan-blog/service/upload/qiniu"
	"mofan-blog/utils"
)

type Upload interface {
	UploadFile(file multipart.File, fileHandler *multipart.FileHeader) (string, int)
}

var Upup Upload

func InitUpload() {
	switch utils.Plugin {
	case "gsky":
		Upup = gsky.Gsky{}
	case "qiniu":
		Upup = qiniu.QiNiu{}
	}
}
