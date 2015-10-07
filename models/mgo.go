package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session
var momentDb *mgo.Database

func getSession() *mgo.Session {
	if session == nil {
		sess, err := mgo.Dial(beego.AppConfig.String("mongo_url"))
		beego.Debug("dial mongodb")
		if err != nil {
			panic(err)
		}
		session = sess
	}
	return session
}

func GetMomentDb() *mgo.Database {
	if momentDb == nil {
		beego.Debug("get moment db")
		momentDb = getSession().DB(beego.AppConfig.String("moments_db"))
	}
	return momentDb
}

func CloseSession() {
	beego.Debug("close session")
	if session != nil {
		session.Close()
	}
}
