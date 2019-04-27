package routing

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
		"dos":   ctx.DefaultQuery("dos", "default-dos"),
		"query": ctx.Request.URL.Query(),
		"now":   time.Now().Format(TIME_FORMAT),
		"uno":   ctx.Query("uno"),
	}
	ctx.JSON(http.StatusOK, response)
}

func GetParams(ctx *gin.Context) {
	response := gin.H{
		"dos":    ctx.Param("dos"),
		"params": ctx.Params,
		"now":    time.Now().Format(TIME_FORMAT),
		"tres":   ctx.Param("tres"),
	}
	ctx.JSON(http.StatusOK, response)
}
