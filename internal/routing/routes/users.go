package routes // import "github.com/jacekk/go-rest-api-playground/internal/routing/routes"

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"github.com/jacekk/go-rest-api-playground/internal/database"
)

func GetUsers(ctx *gin.Context) {
	offset, limit := GetPaginationFromQuery(ctx)
	entities, err := database.GetUsers(offset, limit)

	if err != nil {
		ctx.String(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, entities)
}

func GetUser(ctx *gin.Context) {
	rawId := ctx.Param("id")
	id, err := strconv.ParseInt(rawId, 10, 64)

	if err != nil {
		msg := fmt.Sprintf("Id '%s' is NOT valid.", rawId)
		ctx.String(http.StatusBadRequest, msg)
		return
	}

	entity, err := database.GetUser(id)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, entity)
}

func CreateUser(ctx *gin.Context) {
	var user database.UserAccount
	ctx.BindJSON(&user)
	validation := validate.Struct(user)

	if !validation.Validate() {
		ctx.JSON(http.StatusUnprocessableEntity, validation.Errors)
		return
	}

	entity, err := database.CreateUser(user)

	if err != nil {
		ctx.String(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, entity)
}

func DeleteUser(ctx *gin.Context) {
	rawId := ctx.Param("id")
	id, err := strconv.ParseInt(rawId, 10, 64)

	if err != nil {
		msg := fmt.Sprintf("Id '%s' is NOT valid.", rawId)
		ctx.String(http.StatusUnprocessableEntity, msg)
		return
	}

	err = database.DeleteUserById(id)

	if err != nil {
		ctx.String(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)
}
