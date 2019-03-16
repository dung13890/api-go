package validators

import (
	"fmt"

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

func UserStore(m *models.User) []string {
	validate := validator.New()
	validate.RegisterValidation("unique", validUnique)
	validations := []string{}
	err := validate.Struct(m)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			validations = append(validations, err.Error())
			return validations
		}

		for _, e := range err.(validator.ValidationErrors) {
			validations = append(validations, fmt.Sprintf("%s: Rule %s is invalid", e.Field(), e.Tag()))
		}
	}

	return validations
}

func validUnique(fl validator.FieldLevel) bool {
	currentField, _, _ := fl.GetStructFieldOK()
	value := fl.Field().String() // value
	fmt.Println(currentField, value)
	return false
}
