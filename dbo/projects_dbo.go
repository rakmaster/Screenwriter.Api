package dbo

import (
	. "../models"
	"gopkg.in/mgo.v2/bson"
)

func FindAllProjects() ([]Project, error) {
	var projects []Project
	err := db.C("projects").Find(bson.M{}).All(&projects)
	return projects, err
}

func InsertProject(project Project) error {
	err := db.C("projects").Insert(&project)
	return err
}
