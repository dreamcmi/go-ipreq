package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
)

var IpRe= ipReApi{}

type ipReApi struct {}

// Index is a demonstration route handler for output Real IP
func (*ipReApi) Index(r *ghttp.Request) {
	// 返回ip
	r.Response.Write(gstr.StrTillEx(r.RemoteAddr, string(':')))
}

func (*ipReApi) Port(r *ghttp.Request)  {
	// 返回端口
	r.Response.Write(gstr.StrEx(r.RemoteAddr,string(':')))
}

func (*ipReApi) Url(r *ghttp.Request)  {
	// 返回url
	r.Response.Write(r.RemoteAddr)
}