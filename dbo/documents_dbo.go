package dbo

import (
	. "../models"
	"gopkg.in/mgo.v2/bson"
)

func FindAllDocuments() ([]Document, error) {
	var documents []Document
	err := db.C("documents").Find(bson.M{}).All(&documents)
	return documents, err
}
