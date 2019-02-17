package repositories

import (
	"fmt"

	"github.com/dung13890/api-go/app/contracts"
	"github.com/dung13890/api-go/models"
)

type UserMongo struct {
}

func UserMongoImplement() contracts.UserRepo {
	return &UserMongo{}
}

func (u *UserMongo) Login(p models.User) error {
	fmt.Println(p.UserName)

	return nil
}
