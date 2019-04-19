package routing // import "github.com/jacekk/go-rest-api-playground/internal/routing"

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"github.com/jacekk/go-rest-api-playground/internal/database"
	"github.com/raja/argon2pw"
)

type LoginRequest struct {
	Email    string `validate:"required|email"`
	Password string `validate:"required"`
}
type LoginResponse struct {
	IsLoggedIn bool
}

func isPassValid(request LoginRequest) bool {
	entity, _ := database.GetUserByEmail(request.Email)

	if entity == nil {
		return false
	}

	isValid, err := argon2pw.CompareHashWithPassword(entity.PasswordHash, request.Password)

	if err != nil {
		return false
	}

	return isValid
}

func Login(ctx *gin.Context) {
	var request LoginRequest
	ctx.ShouldBindJSON(&request)
	validation := validate.Struct(request)

	if !validation.Validate() {
		ctx.JSON(http.StatusUnprocessableEntity, validation.Errors)
		return
	}

	var response LoginResponse
	status := http.StatusForbidden
	isValid := isPassValid(request)
	response.IsLoggedIn = isValid

	if isValid {
		status = http.StatusAccepted
	}

	ctx.JSON(status, response)
}
