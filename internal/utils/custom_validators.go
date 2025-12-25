package utils

import "github.com/go-playground/validator/v10"

func StrongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	hasLetter := false
	hasNumber := false

	for _, c := range password {
		switch {
		case c >= 'a' && c <= 'z', c >= 'A' && c <= 'Z':
			hasLetter = true
		case c >= '0' && c <= '9':
			hasNumber = true
		}
	}

	return hasLetter && hasNumber
}
