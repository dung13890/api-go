package contracts

import "github.com/dung13890/api-go/models"

type UserRepo interface {
	Find(id string) (models.User, error)
	FindByEmail(email string) (models.User, error)
	Store(p models.User) (models.User, error)
	GetAll() ([]models.User, error)
}
