package routes // import "github.com/jacekk/go-rest-api-playground/internal/routing/routes"

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"github.com/jacekk/go-rest-api-playground/internal/database"
	"github.com/raja/argon2pw"
)

type LoginRequest struct {
	Name     string `validate:"required"`
	Password string `validate:"required"`
}
type LoginResponse struct {
	IsLoggedIn bool
}

func isPassValid(request LoginRequest) bool {
	entity, _ := database.GetUserByName(request.Name)

	if entity == nil {
		return false
	}

	isValid, err := argon2pw.CompareHashWithPassword(entity.Password, request.Password)

	if err != nil {
		return false
	}

	return isValid
}

func Login(ctx *gin.Context) {
	var request LoginRequest
	var response LoginResponse
	ctx.BindJSON(&request)
	validation := validate.Struct(request)
	status := http.StatusForbidden

	if validation.Validate() {
		isValid := isPassValid(request)
		response.IsLoggedIn = isValid
		if isValid {
			status = http.StatusAccepted
		}
	}

	ctx.JSON(status, response)
}
