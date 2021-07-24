package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	Plugin      string
	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string

	GskyServer string
	BaseUrl    string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件错误")
	}
	LoadServer(file)
	LoadData(file)
	LoadUpload(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("123456789")

}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("root")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func LoadUpload(file *ini.File) {
	Plugin = file.Section("upload").Key("Plugin").MustString("")

	GskyServer = file.Section("upload").Key("GskyServer").MustString("")
	BaseUrl = file.Section("upload").Key("BaseUrl").MustString("")

	AccessKey = file.Section("upload").Key("AccessKey").MustString("")
	SecretKey = file.Section("upload").Key("SecretKey").MustString("")
	Bucket = file.Section("upload").Key("Bucket").MustString("")
	QiniuServer = file.Section("upload").Key("QiniuServer").MustString("")
}
