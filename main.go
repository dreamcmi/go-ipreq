package main

import (
	_ "ipReq-http/boot"
	_ "ipReq-http/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
