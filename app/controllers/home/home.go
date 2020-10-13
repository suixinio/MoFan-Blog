package home

import (
	"mofan-blog/app/service/home"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

// Welcome path " / "
func Welcome(r *ghttp.Request) {
	err := r.Response.WriteTpl("home/welcome.html", g.Map{})

	if err != nil {
		glog.Error(err)
	}
}

//Index path:/index
func Index(r *ghttp.Request) {

	limit := g.Config().GetInt("limit")

	page := 1

	document_list := home.List(limit, page)

	err := r.Response.WriteTpl("home/index.html", g.Map{"list": document_list})

	if err != nil {
		glog.Error(err)
	}
}
