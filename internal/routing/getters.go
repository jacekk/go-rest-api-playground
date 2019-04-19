package routing // import "github.com/jacekk/go-rest-api-playground/internal/routing"

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPlainText(ctx *gin.Context) {
	response := fmt.Sprintf("Works now --> %s", time.Now())
	ctx.String(http.StatusOK, response)
}

func GetQuery(ctx *gin.Context) {
	response := gin.H{
		"now":   time.Now(),
		"uno":   ctx.Query("uno"),
		"dos":   ctx.DefaultQuery("dos", "default-dos"),
		"query": ctx.Request.URL.Query(),
	}
	ctx.JSON(http.StatusOK, response)
}

func GetParams(ctx *gin.Context) {
	response := gin.H{
		"now":    time.Now(),
		"dos":    ctx.Param("dos"),
		"tres":   ctx.Param("tres"),
		"params": ctx.Params,
	}
	ctx.JSON(http.StatusOK, response)
}
