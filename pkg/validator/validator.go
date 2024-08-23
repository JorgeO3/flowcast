package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validator represent the validator contract
type Validator interface {
	Validate(i interface{}) error
}

type goPlaygroundValidator struct {
	validator *validator.Validate
}

// Validate implements Validator.
func (g *goPlaygroundValidator) Validate(i interface{}) error {
	return g.validator.Struct(i)
}

// NewGoPlaygroundValidator create a new instance of Validator
func New() Validator {
	validator := validator.New()

	validator.RegisterValidation("fullname", validateFullName)

	return &goPlaygroundValidator{validator}
}

func validateFullName(fl validator.FieldLevel) bool {
	fullName := fl.Field().String()
	parts := strings.Split(fullName, " ")
	return len(parts) >= 4
}
