package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:BlockController"] = append(beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:BlockController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:BlockController"] = append(beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:BlockController"],
		beego.ControllerComments{
			"Delete",
			`/`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:CommentController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:CommentController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:CommentController"],
		beego.ControllerComments{
			"Delete",
			`/`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:LikeController"] = append(beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:LikeController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:LikeController"] = append(beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:LikeController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:LikeController"] = append(beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:LikeController"],
		beego.ControllerComments{
			"Delete",
			`/`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:MomentController"] = append(beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:MomentController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:MomentController"] = append(beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:MomentController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:MomentController"] = append(beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:MomentController"],
		beego.ControllerComments{
			"Delete",
			`/`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:MomentController"] = append(beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:MomentController"],
		beego.ControllerComments{
			"Pull",
			`/pull`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:MomentController"] = append(beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:MomentController"],
		beego.ControllerComments{
			"Exist",
			`/exist`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:UnshareController"] = append(beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:UnshareController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:UnshareController"] = append(beego.GlobalControllerRouter["github.com/soyking/MomentAPI/controllers:UnshareController"],
		beego.ControllerComments{
			"Delete",
			`/`,
			[]string{"delete"},
			nil})

}
