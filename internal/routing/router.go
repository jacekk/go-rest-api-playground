package routing // import "github.com/jacekk/go-rest-api-playground/internal/routing"

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jacekk/go-rest-api-playground/internal/routing/routes"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", routes.PingRoute)

	return router
}

func InitRouter(port string) {
	router := setupRouter()
	router.Run(fmt.Sprintf(":%s", port))
}
