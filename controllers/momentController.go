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
		result.Result = err.Error()
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
			m.Data["json"] = errorInfo
		} else {
			m.Data["json"] = successInfo
		}
	} else {
		beego.Debug(err)
		m.Data["json"] = errorInfo
	}
	m.ServeJson()
}

type MomentDeleteInfo struct {
	MomentId string `json:"MomentId"`
	UserId   string `json:"UserId"`
}

func (m *MomentController) Delete() {
	var ob MomentDeleteInfo
	err := json.Unmarshal(m.Ctx.Input.RequestBody, &ob)
	if err == nil {
		err := models.DeleteMoment(ob.MomentId, ob.UserId)
		if err != nil {
			beego.Debug(err)
			m.Data["json"] = errorInfo
		} else {
			m.Data["json"] = successInfo
		}
	} else {
		beego.Debug(err)
		m.Data["json"] = errorInfo
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

type MomentExistController struct {
	beego.Controller
}

func (m *MomentExistController) Get() {
	momentId := m.GetString("MomentId")
	if momentId != "" {
		exist, err := models.MomentExist(momentId)
		if err != nil {
			beego.Debug(err)
			m.Data["json"] = errorInfo
		} else {
			if exist {
				m.Data["json"] = "{Result:exist}"
			} else {
				m.Data["json"] = "{Result:not exist}"
			}
		}
	} else {
		m.Data["json"] = errorInfo
	}
	m.ServeJson()
}
