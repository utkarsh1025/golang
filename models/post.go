package models

import "gopkg.in/mgo.v2/bson"


type Post struct {
	Id       bson.ObjectId       `json:"id" bson:"_id"`
	Caption  string              `json:"caption" bson:"caption"`
	Imageurl string              `json:"imageurl" bson:"imageurl"`
	Time     bson.MongoTimestamp        `json:"time" bson:"time"`
	Userid   string     `json:"userid" bson:"userid"`
}


