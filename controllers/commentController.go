package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/soyking/MomentAPI/models"
)

type CommentResult struct {
	Result   string
	Comments []models.Comment
}

type CommentController struct {
	beego.Controller
}

func (c *CommentController) Get() {
	momentId := c.GetString("MomentId")
	userId := c.GetString("UserId")
	result := CommentResult{}
	if userId != "" && momentId != "" {
		ob, err := models.GetComments(momentId, userId)
		if err != nil {
			beego.Debug(err)
			result.Result = "error"
		} else {
			result.Result = "success"
			result.Comments = ob
		}
	} else {
		result.Result = "error"
	}
	c.Data["json"] = result
	c.ServeJson()
}

func (c *CommentController) Post() {
	var ob models.Comment
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	if err == nil {
		err := models.SaveComment(ob)
		if err != nil {
			beego.Debug(err)
			c.Data["json"] = successInfo
		} else {
			c.Data["json"] = errorInfo
		}
	} else {
		beego.Debug(err)
		c.Data["json"] = errorInfo
	}
	c.ServeJson()
}

type CommentDeleteInfo struct {
	CommentId string `json:"CommentId"`
	UserId    string `json:"UserId"`
}

func (c *CommentController) Delete() {
	var ob CommentDeleteInfo
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	if err == nil {
		err := models.DeleteCommentByCommentId(ob.CommentId, ob.UserId)
		if err != nil {
			beego.Debug(err)
			c.Data["json"] = errorInfo
		} else {
			c.Data["json"] = successInfo
		}
	} else {
		beego.Debug(err)
		c.Data["json"] = errorInfo
	}
	c.ServeJson()
}
