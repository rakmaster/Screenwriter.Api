package swagger

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProjectsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "projects"
)

// Find list of projects
func (p *ProjectsDAO) FindAll() ([]Project, error) {
	var projects []Project
	err := db.C(COLLECTION).Find(bson.M{}).All(&projects)
	return projects, err
}

// Find a project by its id
func (p *ProjectsDAO) FindById(id string) (Project, error) {
	var project Project
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&project)
	return project, err
}

// Insert a project into database
func (p *ProjectsDAO) Insert(project Project) error {
	err := db.C(COLLECTION).Insert(&project)
	return err
}

// Delete an existing project
func (p *ProjectsDAO) Delete(project Project) error {
	err := db.C(COLLECTION).Remove(&project)
	return err
}

// Update an existing project
func (p *ProjectsDAO) Update(project Project) error {
	err := db.C(COLLECTION).UpdateId(project.ID, &project)
	return err
}
