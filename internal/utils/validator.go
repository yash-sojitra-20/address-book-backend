package utils

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()

	// Register custom validators
	Validate.RegisterValidation("strong_password", StrongPassword)
}
