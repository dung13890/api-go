package config

import (
	"fmt"
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
)

type Session struct {
	session *mgo.Session
}

func NewSession() *Session {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{fmt.Sprintf("%s:%s", GetString("mongo.host"), GetString("mongo.port"))},
		Username: GetString("mongo.user"),
		Password: GetString("mongo.pass"),
		Timeout:  time.Duration(GetInt("mongo.timeout")) * time.Second,
	})

	if err != nil {
		log.Fatalf("[ConnectDB]: %s\n", err)
	}
	session.SetMode(mgo.Monotonic, true)

	return &Session{session}
}

func (s *Session) Copy() *mgo.Session {
	return s.session.Copy()
}

func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}
