package routes // import "github.com/jacekk/go-rest-api-playground/internal/routing/routes"

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PostXml(ctx *gin.Context) {
	ctx.XML(http.StatusOK, gin.H{
		"now":   time.Now(),
		"query": fmt.Sprintf("%v", ctx.Request.URL.Query()),
	})
}

func PostYml(ctx *gin.Context) {
	ctx.YAML(http.StatusOK, gin.H{
		"now":   time.Now(),
		"query": ctx.Request.URL.Query(),
	})
}

type DosData struct {
	Bar             string // passed lowercased works as well
	FooInternalName string `json:"foo"`
}

func PostJson(ctx *gin.Context) {
	var dos DosData
	err := ctx.ShouldBindJSON(&dos)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var reponse struct {
		Uno  string `json:"now"`
		Dos  DosData
		Tres int64 `json:"tres"`
	}

	tres, _ := strconv.ParseInt(dos.Bar, 10, 64)

	reponse.Uno = fmt.Sprint(time.Now())
	reponse.Dos = dos
	reponse.Tres = tres

	ctx.JSON(http.StatusOK, reponse)
}
