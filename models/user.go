package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Password string        `json:"password" bson:"password"`
	Email    string        `json:"email" bson:"email"`
	
}



