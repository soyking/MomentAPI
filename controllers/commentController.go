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
	beego.Debug(ob)
	if err == nil {
		err := models.SaveComment(ob)
		if err != nil {
			beego.Debug(err)
			c.Data["json"] = "{Result:error}"
		} else {
			c.Data["json"] = "{Result:success}"
		}
	} else {
		beego.Debug(err)
		c.Data["json"] = "{Result:error}"
	}
	c.ServeJson()
}

type CommentDeleteController struct {
	beego.Controller
}

func (c *CommentDeleteController) Post() {
	CommentId := c.GetString("CommentId")
	UserId := c.GetString("UserId")
	if CommentId != "" && UserId != "" {
		err := models.DeleteCommentByCommentId(CommentId, UserId)
		if err != nil {
			beego.Debug(err)
			c.Data["json"] = "{Result:error}"
		} else {
			c.Data["json"] = "{Result:success}"
		}
	} else {
		c.Data["json"] = "{Result:error}"
	}
	c.ServeJson()
}
