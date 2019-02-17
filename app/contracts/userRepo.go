package contracts

import "github.com/dung13890/api-go/models"

type UserRepo interface {
	Login(p models.User) error
}
