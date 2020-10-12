package router

import (
	"mofan-blog/app/controllers/home"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

func bindRouter() {

	s := g.Server()

	//Home
	s.BindHandler("/", home.Welcome)
	//Post list
	s.BindHandler("/index", home.Index)

}

func init() {
	glog.Info("########router start...")
	s := g.Server()
	bindRouter()
	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	s.SetRewrite("/favicon.ico", "/res/images/favicon.ico")

	glog.Info("########router finish.")
}
