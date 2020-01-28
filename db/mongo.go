package db

import (
	"gopkg.in/mgo.v2"
	"rbq-be/config"
	"rbq-be/utils"
)

var (
	Session    *mgo.Session
	Connection *mgo.DialInfo
)

func Connect() {
	url := config.GetConfig().MongoContentUrl
	mongo, err := mgo.ParseURL(url)
	utils.Check(err)
	session, err1 := mgo.Dial(url)
	utils.Check(err1)
	Session = session
	Connection = mongo
}
