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

// @Title get
// @Description get likes by MomentId and UserId
// @Param	MomentId		query 	string	true		"The moment you want to query"
// @Param	UserId		query 	string	true		"Your UserId"
// @Success 200 {object} controllers.LikeResult
// @Failure 200 {"Result": "error","Likes": null}
// @router / [get]
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

// @Title post
// @Description like
// @Param	body		body 	models.Like	true		"The like content"
// @Success 200 {Result:success}
// @Failure 200 {Result:error}
// @router / [post]
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

// @Title cancel
// @Description cancel the like
// @Param	body		body 	controllers.LikeDeleteInfo	true		"The like you want to cancel"
// @Success 200 {Result:success}
// @Failure 200 {Result:error}
// @router / [delete]
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
