package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/soyking/MomentAPI/models"
)

type UserRelationInfo struct {
	FollowId   string `json:"FollowId"`
	FollowedId string `json:"FollowedId"`
}

type BlockController struct {
	beego.Controller
}

// @Title block
// @Description block someone
// @Param	body		body 	models.UserRelation 	true		"FollowId block FollowedId"
// @Success 200 {Result:success}
// @Failure 200 {Result:error}
// @router / [post]
func (b *BlockController) Post() {
	var ob UserRelationInfo
	err := json.Unmarshal(b.Ctx.Input.RequestBody, &ob)
	if err == nil {
		err := models.BlockAction(ob.FollowId, ob.FollowedId, true)
		if err != nil {
			beego.Debug(err)
			b.Data["json"] = errorInfo
		} else {
			b.Data["json"] = successInfo
		}
	} else {
		beego.Debug(err)
		b.Data["json"] = errorInfo
	}
	b.ServeJson()
}

// @Title unblock
// @Description cancel block someone
// @Param	body		body 	models.UserRelation 	true		"FollowId cancel block FollowedId"
// @Success 200 {Result:success}
// @Failure 200 {Result:error}
// @router / [delete]
func (b *BlockController) Delete() {
	var ob UserRelationInfo
	err := json.Unmarshal(b.Ctx.Input.RequestBody, &ob)
	if err == nil {
		err := models.BlockAction(ob.FollowId, ob.FollowedId, false)
		if err != nil {
			beego.Debug(err)
			b.Data["json"] = errorInfo
		} else {
			b.Data["json"] = successInfo
		}
	} else {
		beego.Debug(err)
		b.Data["json"] = errorInfo
	}
	b.ServeJson()
}

type UnshareController struct {
	beego.Controller
}

// @Title unshare
// @Description unshare with someone
// @Param	body		body 	models.UserRelation 	true		"FollowId unshare with FollowedId"
// @Success 200 {Result:success}
// @Failure 200 {Result:error}
// @router / [post]
func (u *UnshareController) Post() {
	var ob UserRelationInfo
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err == nil {
		err := models.UnshareAction(ob.FollowId, ob.FollowedId, true)
		if err != nil {
			beego.Debug(err)
			u.Data["json"] = errorInfo
		} else {
			u.Data["json"] = successInfo
		}
	} else {
		beego.Debug(err)
		u.Data["json"] = errorInfo
	}
	u.ServeJson()
}

// @Title share
// @Description cancel unshare someone
// @Param	body		body 	models.UserRelation 	true		"FollowId cancel unshare FollowedId"
// @Success 200 {Result:success}
// @Failure 200 {Result:error}
// @router / [delete]
func (u *UnshareController) Delete() {
	var ob UserRelationInfo
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err == nil {
		err := models.UnshareAction(ob.FollowId, ob.FollowedId, false)
		if err != nil {
			beego.Debug(err)
			u.Data["json"] = errorInfo
		} else {
			u.Data["json"] = successInfo
		}
	} else {
		beego.Debug(err)
		u.Data["json"] = errorInfo
	}
	u.ServeJson()
}
