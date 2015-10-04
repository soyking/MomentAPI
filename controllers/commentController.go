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

// @Title get
// @Description get comments by MomentId and UserId
// @Param	MomentId		query 	string	true		"The moment you want to query"
// @Param	UserId		query 	string	true		"Your UserId"
// @Success 200 {object} controllers.CommentResult
// @Failure 200 {"Result": "error","Comments": null}
// @router / [get]
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

// @Title publish
// @Description publish comment
// @Param	body		body 	models.Comment	true		"The comment content"
// @Success 200 {Result:success}
// @Failure 200 {Result:error}
// @router / [post]
func (c *CommentController) Post() {
	var ob models.Comment
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	if err == nil {
		err := models.SaveComment(ob)
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

type CommentDeleteInfo struct {
	CommentId string `json:"CommentId"`
	UserId    string `json:"UserId"`
}

// @Title delete
// @Description delete the comment
// @Param	body		body 	controllers.CommentDeleteInfo	true		"The comment you want to delete"
// @Success 200 {Result:success}
// @Failure 200 {Result:error}
// @router / [delete]
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
