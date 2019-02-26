package repositories

import (
	"github.com/dung13890/api-go/app/contracts"
	"github.com/dung13890/api-go/models"
)

type UserMongo struct {
}

func UserMongoImplement() contracts.UserRepo {
	return &UserMongo{}
}

func (u *UserMongo) Login(p models.User) models.User {

	return p
}
