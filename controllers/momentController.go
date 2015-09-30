package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/soyking/MomentAPI/models"
)

type MomentResult struct {
	Result  string
	Moments []models.Moment
}

type MomentController struct {
	beego.Controller
}

func (m *MomentController) Get() {
	userId := m.GetString("UserId")
	timestamp, err := m.GetInt64("Timestamp")
	result := MomentResult{}
	if err == nil && userId != "" {
		ob, err := models.GetMomentsByUserId(userId, timestamp)
		if err != nil {
			beego.Debug(err)
			result.Result = "error"
		} else {
			result.Result = "success"
			result.Moments = ob
		}
	} else {
		beego.Debug(err)
		result.Result = "error"
	}
	m.Data["json"] = result
	m.ServeJson()
}

func (m *MomentController) Post() {
	var ob models.Moment
	err := json.Unmarshal(m.Ctx.Input.RequestBody, &ob)
	if err == nil {
		err := models.SaveMoment(ob)
		if err != nil {
			beego.Debug(err)
			m.Data["json"] = "{Result:error}"
		} else {
			m.Data["json"] = "{Result:success}"
		}
	} else {
		beego.Debug(err)
		m.Data["json"] = "{Result:error}"
	}
	m.ServeJson()
}

type MomentDeleteController struct {
	beego.Controller
}

func (m *MomentDeleteController) Post() {
	momentId := m.GetString("MomentId")
	userId := m.GetString("UserId")
	if momentId != "" && userId != "" {
		err := models.DeleteMoment(momentId, userId)
		if err != nil {
			beego.Debug(err)
			m.Data["json"] = "{Result:error}"
		} else {
			m.Data["json"] = "{Result:success}"
		}
	} else {
		m.Data["json"] = "{Result:error}"
	}
	m.ServeJson()
}

type MomentPullController struct {
	beego.Controller
}

func (m *MomentPullController) Get() {
	userId := m.GetString("UserId")
	timestamp, err := m.GetInt64("Timestamp")
	result := MomentResult{}
	if err == nil && userId != "" {
		ob, err := models.GetMomentsByTimestamp(userId, timestamp)
		if err != nil {
			beego.Debug(err)
			result.Result = "error"
		} else {
			result.Result = "success"
			result.Moments = ob
		}
	} else {
		beego.Debug(err)
		result.Result = err.Error()
	}
	m.Data["json"] = result
	m.ServeJson()
}
