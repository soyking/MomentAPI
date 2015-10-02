package main

import (
	"github.com/astaxie/beego"
	"github.com/soyking/MomentAPI/controllers"
)

func main() {
	// user relation operations
	beego.Router("/block", &controllers.BlockController{})
	beego.Router("/unshare", &controllers.UnshareController{})

	// moment operations
	beego.Router("/moment", &controllers.MomentController{})
	beego.Router("/moment/pull", &controllers.MomentPullController{})
	beego.Router("/moment/exist", &controllers.MomentExistController{})

	// like operations
	beego.Router("/like", &controllers.LikeController{})

	// comment operations
	beego.Router("/comment", &controllers.CommentController{})

	beego.Run()
}
