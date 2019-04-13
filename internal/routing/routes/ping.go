package routes // import "github.com/jacekk/go-rest-api-playground/internal/routing/routes"

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func PingRoute(ctx *gin.Context) {
	response := fmt.Sprintf("pong --> %s", time.Now())
	ctx.String(http.StatusOK, response)
}
