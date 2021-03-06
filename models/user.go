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
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id         bson.ObjectId `bson:"_id" json:"id,omitempty"`
	Username   string        `bson:"username" json:"username,omitempty"`
	FirstName  string        `bson:"first_name" json:"firstName,omitempty"`
	LastName   string        `bson:"last_name" json:"lastName,omitempty"`
	Email      string        `bson:"email" json:"email,omitempty"`
	Password   string        `bson:"passsword" json:"password,omitempty"`
	Phone      string        `bson:"phone" json:"phone,omitempty"`
	UserStatus int32         `bson:"user_status" json:"userStatus,omitempty"`
}
