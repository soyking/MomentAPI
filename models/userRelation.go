package models

import (
	// "github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRelation struct {
	FollowId   string `bson:"FollowId"`
	FollowedId string `bson:"FollowedId"`
	Block      bool   `bson:"Block"`
	Unshare    bool   `bson:"Unshare"`
}

var user_c *mgo.Collection

func init() {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	user_c = session.DB(db).C("UserRelation")
}

func GetRealFriends(userId string) (friends []string, err error) {
	relations := []UserRelation{}
	err = user_c.Find(bson.M{"FollowId": userId, "Block": false}).All(&relations)
	if err != nil {
		return
	}

	// including self
	friends = append(friends, userId)
	relation := UserRelation{}
	for _, v := range relations {
		friendId := v.FollowedId
		err = user_c.Find(bson.M{"FollowId": friendId, "FollowedId": userId}).One(&relation)
		if err != nil {
			return
		}
		if relation.Unshare == false {
			friends = append(friends, friendId)
		}
	}
	return
}

func GetAllFriends(userId string) (friends []string, err error) {
	relations := []UserRelation{}
	err = user_c.Find(bson.M{"FollowId": userId}).All(&relations)
	if err != nil {
		return
	}

	friends = append(friends, userId)
	for _, v := range relations {
		friends = append(friends, v.FollowedId)
	}
	return
}

func BlockAction(followId string, followedId string, block bool) (err error) {
	err = user_c.Update(bson.M{"FollowId": followId, "FollowedId": followedId},
		bson.M{"$set": bson.M{"Block": block}})
	return
}

func UnshareAction(followId string, followedId string, unshare bool) (err error) {
	err = user_c.Update(bson.M{"FollowId": followId, "FollowedId": followedId},
		bson.M{"$set": bson.M{"Unshare": unshare}})
	return
}
