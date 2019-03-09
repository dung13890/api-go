package validators

import (
	"fmt"

	validator "gopkg.in/go-playground/validator.v9"
)

func renderValidations(validations *[]string, field string, e error) {
	if e != nil {
		if vs, ok := e.(validator.ValidationErrors); ok {
			for _, v := range vs {
				*validations = append(*validations, fmt.Sprintf("%s: Rule %s is invalid", field, v.Tag()))
			}
		}
	}
}
