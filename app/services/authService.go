package services

import (
	"context"
	"time"

	"github.com/dung13890/api-go/app/contracts"
	"github.com/dung13890/api-go/models"
)

type AuthService struct {
	userRepo contracts.UserRepo
}

func NewAuthService(a contracts.UserRepo) AuthService {
	return AuthService{
		userRepo: a,
	}
}

func (a AuthService) Login(c context.Context, m models.User) models.User {
	_, cancel := context.WithTimeout(c, time.Duration(2)*time.Second)
	defer cancel()
	return a.userRepo.Login(m)
}
