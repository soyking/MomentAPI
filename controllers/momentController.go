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

// @Title publish
// @Description publish moment
// @Param	body		body 	models.Moment	true		"The moment content"
// @Success 200 {Result:success}
// @Failure 200 {Result:error}
// @router / [post]
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

// @Title get
// @Description get one's moments by UserId and Timestamp
// @Param	UserId		query 	string	true		"Your UserId"
// @Param	Timestamp		query 	string	true		"Last query time"
// @Success 200 {object} controllers.MomentResult
// @Failure 200 {"Result": "error","Moments": null}
// @router / [get]
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

type MomentDeleteInfo struct {
	MomentId string `json:"MomentId"`
	UserId   string `json:"UserId"`
}

// @Title delete
// @Description delete the moment
// @Param	body		body 	controllers.MomentDeleteInfo	true		"The moment you want to delete"
// @Success 200 {Result:success}
// @Failure 200 {Result:error}
// @router / [delete]
func (m *MomentController) Delete() {
	var ob MomentDeleteInfo
	beego.Debug(string(m.Ctx.Input.RequestBody))
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

// @Title pull
// @Description pull moments by UserId and Timestamp
// @Param	UserId		query 	string	true		"Your UserId"
// @Param	Timestamp		query 	string	true		"Last query time"
// @Success 200 {object} controllers.MomentResult
// @Failure 200 {"Result": "error","Moments": null}
// @router /pull [get]
func (m *MomentController) Pull() {
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

// @Title exist
// @Description find out the moment exist or not
// @Param	MomentId		query 	string	true		"The moment you want to query"
// @Success 200 {Result:(not)exist}
// @Failure 200 {Result:error}
// @router /exist [get]
func (m *MomentController) Exist() {
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
