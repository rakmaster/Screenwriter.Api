package dbo

import (
	"encoding/json"
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

func Connect(server string, database string) {
	session, err := mgo.Dial(server)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(database)
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJson(w, code, map[string]string{"error": msg})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
