package helpers // import "github.com/jacekk/go-rest-api-playground/internal/helpers

import (
	"strings"
	"unicode"
)

func IsPasswordStrong(pass string) bool {
	var hasNumber, hasUpper, hasLower, hasSpecial bool
	normalized := strings.TrimSpace(pass)

	for _, char := range normalized {
		switch {
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasNumber && hasUpper && hasLower && hasSpecial
}
