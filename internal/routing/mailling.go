package routing

import (
	"net/http"

	"github.com/gookit/validate"
	"github.com/jacekk/go-rest-api-playground/internal/mailling"

	"github.com/gin-gonic/gin"
)

func MailJson(ctx *gin.Context) {
	var msg mailling.MailMessage
	ctx.ShouldBindJSON(&msg)
	validation := validate.Struct(msg)

	if !validation.Validate() {
		ctx.JSON(http.StatusUnprocessableEntity, validation.Errors)
		return
	}

	err := mailling.SendEmail(msg)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Message sent successfully.",
	})
}
