package providers

import (
	"github.com/dung13890/api-go/app/http/handlers"
	"github.com/dung13890/api-go/app/repositories"
	"github.com/dung13890/api-go/app/services"
	mgo "gopkg.in/mgo.v2"
)

func NewHttpAuthHandle(session *mgo.Session) *handlers.HttpAuthHandler {
	repo := repositories.UserMongoImplement(session)
	service := services.NewAuthService(repo)

	authHandler := &handlers.HttpAuthHandler{
		AuthService: service,
	}

	return authHandler
}
