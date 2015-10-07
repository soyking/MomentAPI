package models

import (
	"errors"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Comment struct {
	CommentId  bson.ObjectId `bson:"_id"`
	MomentId   string        `bson:"MomentId"`
	FromUserId string        `bson:"FromUserId"`
	ToUserId   string        `bson:"ToUserId"`
	Text       string        `bson:"Text"`
	Timestamp  int64         `bson:"Timestamp"`
}

var comment_c *mgo.Collection

func init() {
	comment_c = GetMomentDb().C("Comment")
	index := mgo.Index{
		Key:    []string{"MomentId"},
		Unique: true,
	}
	err := comment_c.EnsureIndex(index)
	if err != nil {
		beego.Debug(err)
	}
}

func SaveComment(c Comment) (err error) {
	// assert the moment exists
	exist, err := MomentExist(c.MomentId)
	if err != nil {
		return
	} else if !exist {
		err = errors.New("moment doesn't exist")
		return
	}

	c.CommentId = bson.NewObjectId()
	c.Timestamp = time.Now().Unix()
	err = comment_c.Insert(&c)
	return
}

func GetComments(momentId string, userId string) (comments []Comment, err error) {
	// user could only see his friends' comments(both in FromUserId and ToUserId)
	friends, err := GetAllFriends(userId)
	if err != nil {
		return
	}

	// including the situation of toUserId=="", assuming that fromUserId won't be ""
	friends = append(friends, "")
	err = comment_c.Find(bson.M{"MomentId": momentId, "FromUserId": bson.M{"$in": friends}, "ToUserId": bson.M{"$in": friends}}).Sort("Timestamp").All(&comments)
	return
}

func DeleteCommentByCommentId(commentId string, userId string) (err error) {
	err = comment_c.Remove(bson.M{"_id": bson.ObjectIdHex(commentId), "FromUserId": userId})
	return
}

func DeleteCommentsByMomentId(momentId string) (err error) {
	_, err = comment_c.RemoveAll(bson.M{"MomentId": momentId})
	return
}
