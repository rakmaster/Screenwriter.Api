/*
 * Screenwriter API
 *
 * Data API for Screenwriter: an online collaborative screenwriting software.
 *
 * API version: 0.0.1
 * Contact: logan@theinitiativepg.ocm
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	dbo "../dbo"
	. "../models"
)

func AddProject(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var project Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		dbo.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	project.Id = bson.NewObjectId()
	if err := dbo.InsertProject(project); err != nil {
		dbo.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	dbo.RespondWithJson(w, http.StatusCreated, project)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetProjectById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetProjectByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetProjectDocuments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := dbo.FindAllProjects()
	if err != nil {
		dbo.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	dbo.RespondWithJson(w, http.StatusOK, projects)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
