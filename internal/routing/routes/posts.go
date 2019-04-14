package routes // import "github.com/jacekk/go-rest-api-playground/internal/routing/routes"

import (
	"github.com/gin-gonic/gin"
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
