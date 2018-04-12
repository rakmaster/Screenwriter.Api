package dbo

import (
	. "../models"
	"gopkg.in/mgo.v2/bson"
)

func FindAllUsers() ([]User, error) {
	var users []User
	err := db.C("users").Find(bson.M{}).All(&users)
	return users, err
}
