package routing // import "github.com/jacekk/go-rest-api-playground/internal/routing"

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"github.com/jacekk/go-rest-api-playground/internal/database"
)

func GetPosts(ctx *gin.Context) {
	offset, limit := GetPaginationFromQuery(ctx)
	posts, err := database.GetPosts(offset, limit)

	if err != nil {
		ctx.String(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

func GetPost(ctx *gin.Context) {
	rawId, id, err := GetIdFromParam(ctx)

	if err != nil {
		ReturnIdError(ctx, rawId)
		return
	}

	post, err := database.GetPost(id)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func CreatePost(ctx *gin.Context) {
	var post database.Post
	ctx.ShouldBindJSON(&post)
	validation := validate.Struct(post)

	if !validation.Validate() {
		ctx.JSON(http.StatusUnprocessableEntity, validation.Errors)
		return
	}

	entity, err := database.CreatePost(post)

	if err != nil {
		ctx.String(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, entity)
}

func DeletePost(ctx *gin.Context) {
	rawId, id, err := GetIdFromParam(ctx)

	if err != nil {
		ReturnIdError(ctx, rawId)
		return
	}

	err = database.DeletePostById(id)

	if err != nil {
		ctx.String(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)
}
