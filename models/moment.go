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
	moment_c = GetMomentDb().C("Moment")
	index := mgo.Index{
		Key: []string{"UserId", "Timestamp"},
	}
	err := moment_c.EnsureIndex(index)
	if err != nil {
		beego.Debug(err)
	}
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
	err = moment_c.Find(bson.M{"UserId": bson.M{"$in": friends}, "Timestamp": bson.M{"$gt": timestamp}}).Sort("-Timestamp").All(&moments)
	return moments, err
}

func DeleteMoment(momentId string, userId string) (err error) {
	err = moment_c.Remove(bson.M{"_id": bson.ObjectIdHex(momentId), "UserId": userId})
	if err != nil {
		return
	}

	err = DeleteLikesByMomentId(momentId)
	if err != nil {
		return
	}

	err = DeleteCommentsByMomentId(momentId)

	return
}

func MomentExist(momentId string) (exist bool, err error) {
	n, err := moment_c.Find(bson.M{"_id": bson.ObjectIdHex(momentId)}).Count()
	exist = (n != 0)
	return
}
