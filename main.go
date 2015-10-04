package main

import (
	_ "github.com/soyking/MomentAPI/docs"
	_ "github.com/soyking/MomentAPI/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
