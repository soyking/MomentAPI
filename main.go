package main

import (
	"github.com/astaxie/beego"
	"github.com/soyking/MomentAPI/controllers"
)

func main() {
	// user relation operations
	beego.Router("/block", &controllers.BlockController{})
	beego.Router("/block/cancel", &controllers.BlockCancelController{})
	beego.Router("/unshare", &controllers.UnshareController{})
	beego.Router("/unshare/cancel", &controllers.UnshareCancelController{})

	// moment operations
	beego.Router("/moment", &controllers.MomentController{})
	beego.Router("/moment/delete", &controllers.MomentDeleteController{})
	beego.Router("/moment/pull", &controllers.MomentPullController{})
	beego.Router("/moment/exist", &controllers.MomentExistController{})

	// like operations
	beego.Router("/like", &controllers.LikeController{})
	beego.Router("/like/cancel", &controllers.LikeCancelController{})

	// comment operations
	beego.Router("/comment", &controllers.CommentController{})
	beego.Router("/comment/delete", &controllers.CommentDeleteController{})

	beego.Run()
}
