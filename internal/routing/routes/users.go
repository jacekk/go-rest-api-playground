package routes // import "github.com/jacekk/go-rest-api-playground/internal/routing/routes"

import (
	"fmt"
	"net/http"

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
	rawId, id, err := GetIdFromParam(ctx)

	if err != nil {
		ReturnIdError(ctx, rawId)
		return
	}

	entity, err := database.GetUser(id)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, entity)
}

func GetUserPosts(ctx *gin.Context) {
	rawId, userId, err := GetIdFromParam(ctx)

	if err != nil {
		ReturnIdError(ctx, rawId)
		return
	}

	user, err := database.GetUser(userId)

	if user == nil {
		msg := fmt.Sprintf("User with ID '%d' was NOT found.", userId)
		ctx.String(http.StatusNotFound, msg)
		return
	}
	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	offset, limit := GetPaginationFromQuery(ctx)
	entities, err := database.GetAuthorPosts(offset, limit, userId)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, entities)
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
	rawId, id, err := GetIdFromParam(ctx)

	if err != nil {
		ReturnIdError(ctx, rawId)
		return
	}

	err = database.DeleteUserById(id)

	if err != nil {
		ctx.String(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)
}
