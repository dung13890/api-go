package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	UserName  string        `json:"username,omitempty" validate:"required,unique"`
	Password  string        `json:"password,omitempty" validate:"required"`
	Email     string        `json:"email,omitempty" validate:"required,email,unique"`
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
