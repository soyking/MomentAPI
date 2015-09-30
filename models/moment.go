package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// OtherData for photo, video, links etc, especially "" for text
// client could handle it by Type info
type Moment struct {
	MomentId  bson.ObjectId `bson:"_id"`
	UserId    string        `bson:"UserId"`
	Type      int           `bson:"Type"`
	Text      string        `bson:"Text"`
	OtherData string        `bson:"OtherData"`
	Source    string        `bson:"Source"`
	Timestamp int64         `bson:"Timestamp"`
}

var moment_c *mgo.Collection

func init() {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	moment_c = session.DB(db).C("Moment")
}

func SaveMoment(m Moment) (err error) {
	m.MomentId = bson.NewObjectId()
	m.Timestamp = time.Now().Unix()
	err = moment_c.Insert(&m)
	return
}

func GetMomentsByUserId(userId string, timestamp int64) (moments []Moment, err error) {
	err = moment_c.Find(bson.M{"UserId": userId, "Timestamp": bson.M{"$gt": timestamp}}).Sort("-Timestamp").All(&moments)
	return
}

func GetMomentsByTimestamp(userId string, timestamp int64) (moments []Moment, err error) {
	friends, err := GetRealFriends(userId)
	if err != nil {
		return
	}
	beego.Debug(friends)
	err = moment_c.Find(bson.M{"UserId": bson.M{"$in": friends}, "Timestamp": bson.M{"$gt": timestamp}}).Sort("-Timestamp").All(&moments)
	return moments, err
}

func DeleteMoment(MomentId string, userId string) (err error) {
	err = moment_c.Remove(bson.M{"_id": bson.ObjectIdHex(MomentId), "UserId": userId})
	if err != nil {
		return
	}

	err = DeleteLikesByMomentId(MomentId)
	if err != nil {
		beego.Debug(err)
		return
	}

	err = DeleteCommentsByMomentId(MomentId)

	return
}
