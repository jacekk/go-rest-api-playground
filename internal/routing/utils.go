package routing // import "github.com/jacekk/go-rest-api-playground/internal/routing"

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReturnIdError(ctx *gin.Context, rawId string) {
	msg := fmt.Sprintf("Id '%s' is NOT valid.", rawId)
	ctx.String(http.StatusBadRequest, msg)
}

func GetIdFromParam(ctx *gin.Context) (string, uint64, error) {
	rawId := ctx.Param("id")
	id, err := strconv.ParseUint(rawId, 10, 64)

	return rawId, id, err
}

func GetPaginationFromQuery(ctx *gin.Context) (uint64, uint64) {
	offsetRaw := ctx.DefaultQuery("offset", strconv.FormatUint(DEFAULT_OFFSET, 10))
	limitRaw := ctx.DefaultQuery("limit", strconv.FormatUint(DEFAULT_LIMIT, 10))
	limit, limitErr := strconv.ParseUint(limitRaw, 10, 64)
	offset, offsetErr := strconv.ParseUint(offsetRaw, 10, 64)

	if limitErr != nil || limit < 1 {
		limit = DEFAULT_LIMIT
	}
	if offsetErr != nil {
		offset = DEFAULT_OFFSET
	}

	return offset, limit
}
