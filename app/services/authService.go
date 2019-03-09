package services

import (
	"context"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dung13890/api-go/app/contracts"
	"github.com/dung13890/api-go/config"
	"github.com/dung13890/api-go/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo contracts.UserRepo
}

func NewAuthService(a contracts.UserRepo) AuthService {
	return AuthService{
		userRepo: a,
	}
}

func (a AuthService) Login(c context.Context, m models.User) (map[string]string, error) {
	_, cancel := context.WithCancel(c)
	defer cancel()

	user, err := a.userRepo.FindByEmail(m.Email)

	if err != nil {
		return map[string]string{}, err
	}

	if checkPasswordHash(m.Password, user.Password) {
		token := jwt.New(jwt.SigningMethodHS256)
		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = user.Id
		claims["exp"] = time.Now().Add(time.Hour * time.Duration(config.GetInt("expire"))).Unix()

		t, errT := token.SignedString([]byte(config.GetString("key")))
		return map[string]string{"token": t}, errT
	}

	return map[string]string{}, errors.New("Errors!")

}

func (a AuthService) Get(c context.Context) ([]models.User, error) {
	_, cancel := context.WithCancel(c)
	defer cancel()

	return a.userRepo.GetAll()
}

func (a AuthService) Info(c context.Context, claims jwt.MapClaims) (models.User, error) {
	_, cancel := context.WithCancel(c)
	defer cancel()

	return a.userRepo.Find(claims["id"].(string))
}

func (a AuthService) Store(c context.Context, p models.User) (models.User, error) {
	_, cancel := context.WithCancel(c)
	defer cancel()

	p.Password, _ = hashPassword(p.Password)

	return a.userRepo.Store(p)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
