package routing // import "github.com/jacekk/go-rest-api-playground/internal/routing"

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jacekk/go-rest-api-playground/internal/routing/routes"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/jacekk/go-rest-api-playground/docs" // docs is generated by Swag CLI, you have to import it.
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/login", routes.Login)
	router.POST("/mail/send", routes.MailJson)
	router.POST("/validate/dics", routes.ValidateDics)
	router.POST("/validate/user", routes.ValidateUser)

	simple := router.Group("/simple")
	{
		simple.GET("/json", routes.GetQuery)
		simple.GET("/plain", routes.GetPlainText)
		simple.GET("/uno/:dos/*tres", routes.GetParams)

		simple.POST("/json", routes.PostJson)
		simple.POST("/xml", routes.PostXml)
		simple.POST("/yml", routes.PostYml)
	}
	posts := router.Group("/posts")
	{
		posts.DELETE("/:id", routes.DeletePost)
		posts.GET("", routes.GetPosts)
		posts.GET("/:id", routes.GetPost)
		posts.POST("", routes.CreatePost)
	}
	users := router.Group("/users")
	{
		users.DELETE("/:id", routes.DeleteUser)
		users.GET("", routes.GetUsers)
		users.GET("/:id", routes.GetUser)
		users.GET("/:id/posts", routes.GetUserPosts)
		users.POST("", routes.CreateUser)
	}

	return router
}

func InitRouter(port string) {
	router := setupRouter()
	err := router.Run(fmt.Sprintf(":%s", port))

	if err != nil {
		panic(err.Error())
	}
}
