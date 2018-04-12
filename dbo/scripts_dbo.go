package dbo

import (
	. "../models"
	"gopkg.in/mgo.v2/bson"
)

func FindAllScripts() ([]Script, error) {
	var scripts []Script
	err := db.C("scripts").Find(bson.M{}).All(&scripts)
	return scripts, err
}
