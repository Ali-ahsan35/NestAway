package main

import (
	_ "smartours/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

