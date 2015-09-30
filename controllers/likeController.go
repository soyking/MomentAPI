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
			result.Result = "error"
		} else {
			result.Result = "success"
			result.Likes = ob
		}
	} else {
		result.Result = "error"
	}
	l.Data["json"] = result
	l.ServeJson()
}

func (l *LikeController) Post() {
	var ob models.Like
	err := json.Unmarshal(l.Ctx.Input.RequestBody, &ob)
	beego.Debug(ob)
	if err == nil {
		err := models.SaveLike(ob)
		if err != nil {
			beego.Debug(err)
			l.Data["json"] = "{Result:error}"
		} else {
			l.Data["json"] = "{Result:success}"
		}
	} else {
		beego.Debug(err)
		l.Data["json"] = "{Result:error}"
	}
	l.ServeJson()
}

type LikeCancelController struct {
	beego.Controller
}

func (l *LikeCancelController) Post() {
	likeId := l.GetString("LikeId")
	userId := l.GetString("UserId")
	if likeId != "" && userId != "" {
		err := models.CancelLike(likeId, userId)
		if err != nil {
			beego.Debug(err)
			l.Data["json"] = "{Result:error}"
		} else {
			l.Data["json"] = "{Result:success}"
		}
	} else {
		l.Data["json"] = "{Result:error}"
	}
	l.ServeJson()
}
