package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	UserName string        `json:"username,omitempty"`
	Password string        `json:"password,omitempty"`
	Email    string        `json:"email,omitempty"`
	Date     time.Time     `json:"date, omitempty"`
}
