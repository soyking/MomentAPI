package controllers

import (
	"github.com/astaxie/beego"
	"github.com/soyking/MomentAPI/models"
)

type BlockController struct {
	beego.Controller
}

func (b *BlockController) Post() {
	followId := b.GetString("FollowId")
	followedId := b.GetString("FollowedId")
	err := models.BlockAction(followId, followedId, true)
	if err != nil {
		beego.Debug(err)
		b.Data["json"] = "{Result:error}"
	} else {
		b.Data["json"] = "{Result:success}"
	}
	b.ServeJson()
}

type BlockCancelController struct {
	beego.Controller
}

func (b *BlockCancelController) Post() {
	followId := b.GetString("FollowId")
	followedId := b.GetString("FollowedId")
	err := models.BlockAction(followId, followedId, false)
	if err != nil {
		beego.Debug(err)
		b.Data["json"] = "{Result:error}"
	} else {
		b.Data["json"] = "{Result:success}"
	}
	b.ServeJson()
}

type UnshareController struct {
	beego.Controller
}

func (u *UnshareController) Post() {
	followId := u.GetString("FollowId")
	followedId := u.GetString("FollowedId")
	err := models.UnshareAction(followId, followedId, true)
	if err != nil {
		beego.Debug(err)
		u.Data["json"] = "{Result:error}"
	} else {
		u.Data["json"] = "{Result:success}"
	}
	u.ServeJson()
}

type UnshareCancelController struct {
	beego.Controller
}

func (u *UnshareCancelController) Post() {
	followId := u.GetString("FollowId")
	followedId := u.GetString("FollowedId")
	err := models.UnshareAction(followId, followedId, false)
	if err != nil {
		beego.Debug(err)
		u.Data["json"] = "{Result:error}"
	} else {
		u.Data["json"] = "{Result:success}"
	}
	u.ServeJson()
}
