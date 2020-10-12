package home

import (
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
	err := r.Response.WriteTpl("home/index.html", g.Map{})

	if err != nil {
		glog.Error(err)
	}
}
