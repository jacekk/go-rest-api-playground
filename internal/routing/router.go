package routing // import "github.com/jacekk/go-rest-api-playground/internal/routing"

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jacekk/go-rest-api-playground/internal/routing/routes"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/validate/user", routes.ValidateUser)

	simple := router.Group("/simple")
	{
		simple.GET("/json", routes.GetQuery)
		simple.GET("/simple/plain", routes.GetPlainText)
		simple.GET("/simple/uno/:dos/*tres", routes.GetParams)

		simple.POST("/simple/json", routes.PostJson)
		simple.POST("/simple/xml", routes.PostXml)
		simple.POST("/simple/yml", routes.PostYml)
	}
	posts := router.Group("/posts")
	{
		posts.DELETE("/:id", routes.DeletePost)
		posts.GET("/", routes.GetPosts)
		posts.POST("/", routes.CreatePost)
	}

	return router
}

func InitRouter(port string) {
	router := setupRouter()
	router.Run(fmt.Sprintf(":%s", port))
}
