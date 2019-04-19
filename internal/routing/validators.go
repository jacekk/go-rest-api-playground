package routing // import "github.com/jacekk/go-rest-api-playground/internal/routing"

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

type UserEntityInValidators struct {
	Name  string `validate:"minLen:6"`
	Age   int    `validate:"required|min:18|max:99"`
	Email string `validate:"required|email"`
}

type DictsValidators struct {
	Uno    int    `validate:"required|in:1,2,3" `
	Dos    int    `validate:"required|enum:4,5,6" `
	Tres   string `validate:"required|in:ONE,TWO,THREE" `
	Cuatro string `validate:"required|enum:four,five,six"`
}

type ValidatorResponse struct {
	Msg    string `json:"message"`
	Entity map[string]string
}

// @Summary Some endpoint summary
// @Description Some endpoint description
// @Tags validators
// @Accept json
// @Produce json
// @todo Param Name body string false "Name of the user (optional)." || causes --> Could not resolve reference because of: Could not resolve pointer: /definitions/string does not exist in document
// @todo Param Age body string true "Age of the user." || causes --> Could not resolve reference because of: Could not resolve pointer: /definitions/int does not exist in document
// @todo Param Email body string true "Email of the user." || causes --> Could not resolve reference because of: Could not resolve pointer: /definitions/string does not exist in document
// @todo Success 200 {object} ValidatorResponse || causes --> Could not resolve reference because of: Could not resolve pointer: /definitions/ValidatorResponse does not exist in document
// @todo Failure 422 {object} ValidatorResponse
// @Router /validate/user [post]
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

// @Summary Another endpoint summary
// @Description Another endpoint description
// @Tags validators
// @Accept json
// @Produce json
// @todo Param
// @Router /validate/dics [post]
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
