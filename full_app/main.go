package main

import (
	_ "full_app/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

