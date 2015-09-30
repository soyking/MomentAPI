package models

import (
	// "github.com/astaxie/beego"
	"errors"
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
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	comment_c = session.DB(db).C("Comment")
}

func SaveComment(c Comment) (err error) {
	// assert the moment exists
	n, err := moment_c.Find(bson.M{"_id": bson.ObjectIdHex(c.MomentId)}).Count()
	if n == 0 {
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
