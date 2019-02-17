package providers

import (
	"github.com/dung13890/api-go/app/http/handlers"
	"github.com/dung13890/api-go/app/repositories"
)

func NewHttpAuthHandle() *handlers.HttpAuthHandler {
	authHandler := &handlers.HttpAuthHandler{
		UserRepo: repositories.UserMongoImplement(),
	}

	return authHandler
}
