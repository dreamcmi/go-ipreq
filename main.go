package main

import (
	"github.com/gogf/gf/frame/g"
	_ "ipReq-http/boot"
	_ "ipReq-http/router"
)

func main() {
	g.Server().Run()
}
