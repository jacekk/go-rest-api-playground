package routing // import "github.com/jacekk/go-rest-api-playground/internal/routing"

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jacekk/go-rest-api-playground/internal/routing/routes"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/simple/json", routes.GetQuery)
	router.GET("/simple/plain", routes.GetPlainText)
	router.GET("/simple/uno/:dos/*tres", routes.GetParams)

	router.POST("/simple/json", routes.PostJson)
	router.POST("/simple/xml", routes.PostXml)
	router.POST("/simple/yml", routes.PostYml)

	return router
}

func InitRouter(port string) {
	router := setupRouter()
	router.Run(fmt.Sprintf(":%s", port))
}
