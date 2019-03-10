package validators

import (
	"github.com/dung13890/api-go/models"
	validator "gopkg.in/go-playground/validator.v9"
)

func Login(m *models.User) []string {
	validate := validator.New()
	validations := []string{}

	errEmail := validate.Var(m.Email, "required,email")
	renderValidations(&validations, "Email", errEmail)
	errPass := validate.Var(m.Password, "required")
	renderValidations(&validations, "Password", errPass)

	return validations
}
