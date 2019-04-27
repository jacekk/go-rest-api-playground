package helpers

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func IsPasswordStrong(pass string) bool {
	const MIN_LEN = 8
	var hasNumber, hasUpper, hasLower, hasSpecial bool
	normalized := strings.TrimSpace(pass)
	passLen := utf8.RuneCountInString(pass)
	hasMinLen := passLen >= MIN_LEN

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

	return hasNumber && hasUpper && hasLower && hasSpecial && hasMinLen
}
