package main

import (
	_ "github.com/soyking/MomentAPI/docs"
	_ "github.com/soyking/MomentAPI/routers"

	"github.com/astaxie/beego"
	"github.com/soyking/MomentAPI/models"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	defer models.CloseSession()
	beego.Run()
}
