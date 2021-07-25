package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	_ "github.com/mattn/go-sqlite3"
	"ipReq-http/app/dao"
)

var IpRe = ipReApi{}

type ipReApi struct{}

// Index is a demonstration route handler for output Real IP
func (*ipReApi) Index(r *ghttp.Request) {
	// 返回ip
	r.Response.Write(gstr.StrTillEx(r.RemoteAddr, string(':')))
	saveIP(r.RemoteAddr)
}

func (*ipReApi) Port(r *ghttp.Request) {
	// 返回端口
	r.Response.Write(gstr.StrEx(r.RemoteAddr, string(':')))
	saveIP(r.RemoteAddr)
}

func (*ipReApi) Url(r *ghttp.Request) {
	// 返回url
	r.Response.Write(r.RemoteAddr)
	saveIP(r.RemoteAddr)
}

func saveIP(url string) {
	ip := gstr.StrTillEx(url, string(':'))
	port := gstr.StrEx(url, string(':'))
	glog.Info(url)
	_, err := g.DB("default").
		Model("user").
		Data(dao.User{UserIP: ip, UserPort: port}).
		Insert()
	if err != nil {
		//panic(err)
		glog.Error(err)
		return
	}
}
