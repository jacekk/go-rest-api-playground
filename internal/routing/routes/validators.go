package routes // import "github.com/jacekk/go-rest-api-playground/internal/routing/routes"

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

type UserEntityInValidators struct {
	Name  string `validate:"required|minLen:6"`
	Age   int    `validate:"required|min:18|max:99"`
	Email string `validate:"required|email"`
}

type DictsValidators struct {
	Uno    int    `validate:"required|in:1,2,3" `
	Dos    int    `validate:"required|enum:4,5,6" `
	Tres   string `validate:"required|in:ONE,TWO,THREE" `
	Cuatro string `validate:"required|enum:four,five,six"`
}

func ValidateUser(ctx *gin.Context) {
	var entity UserEntityInValidators
	ctx.ShouldBindJSON(&entity)
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

func ValidateDics(ctx *gin.Context) {
	var entity DictsValidators
	ctx.ShouldBindJSON(&entity)
	validation := validate.Struct(entity)
	validation.FilterRule("Tres", "trim|upper")
	validation.FilterRule("Cuatro", "trim|lower")

	if !validation.Validate() {
		ctx.JSON(http.StatusUnprocessableEntity, validation.Errors)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":    "Entity is valid.",
		"entity": entity,
	})
}
