package models

import (
	"errors"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Like struct {
	LikeId    bson.ObjectId `bson:"_id"`
	MomentId  string        `bson:"MomentId"`
	UserId    string        `bson:"UserId"`
	Timestamp int64         `bson:"Timestamp"`
}

var like_c *mgo.Collection

func init() {
	like_c = GetMomentDb().C("Like")
	index := mgo.Index{
		Key: []string{"MomentId", "UserId"},
	}
	err := like_c.EnsureIndex(index)
	if err != nil {
		beego.Debug(err)
	}
}

func SaveLike(l Like) (err error) {
	// assert the moment exists
	exist, err := MomentExist(l.MomentId)
	if err != nil {
		return
	} else if !exist {
		err = errors.New("moment doesn't exist")
		return
	}

	//assert the like doesn't exist
	n, _ := like_c.Find(bson.M{"MomentId": l.MomentId, "UserId": l.UserId}).Count()
	if n != 0 {
		err = errors.New("like exists")
		return
	}

	l.LikeId = bson.NewObjectId()
	l.Timestamp = time.Now().Unix()
	err = like_c.Insert(&l)
	return
}

func GetLikes(momentId string, userId string) (likes []Like, err error) {
	// user could only see his friends' likes
	friends, err := GetAllFriends(userId)
	if err != nil {
		return nil, err
	}
	err = like_c.Find(bson.M{"MomentId": momentId, "UserId": bson.M{"$in": friends}}).Sort("Timestamp").All(&likes)
	return
}

func CancelLike(likeId string, userId string) (err error) {
	err = like_c.Remove(bson.M{"_id": bson.ObjectIdHex(likeId), "UserId": userId})
	return
}

func DeleteLikesByMomentId(momentId string) (err error) {
	_, err = like_c.RemoveAll(bson.M{"MomentId": momentId})
	return
}
