package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/soyking/MomentAPI/models"
)

type LikeResult struct {
	Result string
	Likes  []models.Like
}

type LikeController struct {
	beego.Controller
}

func (l *LikeController) Get() {
	momentId := l.GetString("MomentId")
	userId := l.GetString("UserId")
	result := LikeResult{}
	if userId != "" && momentId != "" {
		ob, err := models.GetLikes(momentId, userId)
		if err != nil {
			beego.Debug(err)
			result.Result = errorInfo
		} else {
			result.Result = successInfo
			result.Likes = ob
		}
	} else {
		result.Result = errorInfo
	}
	l.Data["json"] = result
	l.ServeJson()
}

func (l *LikeController) Post() {
	var ob models.Like
	err := json.Unmarshal(l.Ctx.Input.RequestBody, &ob)
	if err == nil {
		err := models.SaveLike(ob)
		if err != nil {
			beego.Debug(err)
			l.Data["json"] = errorInfo
		} else {
			l.Data["json"] = successInfo
		}
	} else {
		beego.Debug(err)
		l.Data["json"] = errorInfo
	}
	l.ServeJson()
}

type LikeDeleteInfo struct {
	LikeId string `json:"LikeId"`
	UserId string `json:"UserId"`
}

func (l *LikeController) Delete() {
	var ob LikeDeleteInfo
	err := json.Unmarshal(l.Ctx.Input.RequestBody, &ob)
	if err == nil {
		err := models.CancelLike(ob.LikeId, ob.UserId)
		if err != nil {
			beego.Debug(err)
			l.Data["json"] = errorInfo
		} else {
			l.Data["json"] = successInfo
		}
	} else {
		beego.Debug(err)
		l.Data["json"] = errorInfo
	}
	l.ServeJson()
}
