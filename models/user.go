package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	UserName  string        `json:"username,omitempty"`
	Password  string        `json:"password,omitempty"`
	Email     string        `json:"email,omitempty"`
	CreatedAt time.Time     `json:"created_at, omitempty"`
	UpdatedAt time.Time     `json:"updated_at, omitempty"`
}

func UserModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"username", "email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}
