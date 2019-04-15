package routes // import "github.com/jacekk/go-rest-api-playground/internal/routing/routes"

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"github.com/jacekk/go-rest-api-playground/internal/database"
	"net/http"
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
