// Package validator provides a validator for structs.
package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// Interface represent the validator contract
type Interface interface {
	Validate(i interface{}) error
}

type goPlaygroundValidator struct {
	validator *validator.Validate
}

// Validate implements Validator.
func (g *goPlaygroundValidator) Validate(i interface{}) error {
	return g.validator.Struct(i)
}

// New create a new instance of Validator
func New() Interface {
	validator := validator.New()

	validator.RegisterValidation("fullname", validateFullName)

	return &goPlaygroundValidator{validator}
}

func validateFullName(fl validator.FieldLevel) bool {
	fullName := fl.Field().String()
	parts := strings.Split(fullName, " ")
	return len(parts) >= 4
}
