package routes // import "github.com/jacekk/go-rest-api-playground/internal/routing/routes"

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"github.com/jacekk/go-rest-api-playground/internal/database"
	"net/http"
	"strconv"
)

func GetPosts(ctx *gin.Context) {
	posts, err := database.GetPosts()

	if err != nil {
		ctx.String(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, posts)
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
	rawId := ctx.Param("id")
	id, err := strconv.ParseInt(rawId, 10, 64)

	if err != nil {
		msg := fmt.Sprintf("Id '%s' is NOT valid.", rawId)
		ctx.String(http.StatusUnprocessableEntity, msg)
		return
	}

	err = database.DeletePostById(id)

	if err != nil {
		ctx.String(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)
}
