package routes // import "github.com/jacekk/go-rest-api-playground/internal/routing/routes"

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

type UserEntityInValidators struct {
	Name  string `validate:"required|minLen:6"`
	Age   int    `validate:"required|int|min:18|max:99"`
	Email string `validate:"required|email"`
}

func ValidateUser(ctx *gin.Context) {
	var entity UserEntityInValidators
	ctx.BindJSON(&entity)
	validation := validate.Struct(entity)

	if !validation.Validate() {
		ctx.JSON(http.StatusUnprocessableEntity, validation.Errors)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":    "Entity is valid.",
		"entity": entity,
	})
}
