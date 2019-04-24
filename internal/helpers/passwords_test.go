package helpers // import "github.com/jacekk/go-rest-api-playground/internal/helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPasswords(t *testing.T) {
	cases := []string{"11aaBB!@", "1qazXSW@3edcVFR$"}

	for _, item := range cases {
		assert.True(t, IsPasswordStrong(item))
	}
}

func TestInvalidPasswords(t *testing.T) {
	cases := []string{
		"",          // empty string
		"1aB!",      // to short
		"aaaBBB!@#", // no digits
		"123BBB!@#", // no lower
		"123aaa!@#", // no upper
		"123aaaBBB", // no specials
		"这是个测试",
		"Спутник и погром",
	}

	for _, item := range cases {
		assert.Falsef(t, IsPasswordStrong(item), "Invalid password failed: %s \n", item)
	}
}
