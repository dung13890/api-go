package providers

import (
	"github.com/dung13890/api-go/app/http/handlers"
	"github.com/dung13890/api-go/app/repositories"
	"github.com/dung13890/api-go/app/services"
)

func NewHttpAuthHandle() *handlers.HttpAuthHandler {
	repo := repositories.UserMongoImplement()
	service := services.NewAuthService(repo)

	authHandler := &handlers.HttpAuthHandler{
		AuthService: service,
	}

	return authHandler
}
