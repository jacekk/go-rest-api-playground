package routes // import "github.com/jacekk/go-rest-api-playground/internal/routing/routes"

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const DEFAULT_LIMIT = 10
const DEFAULT_OFFSET = 0

func GetPaginationFromQuery(ctx *gin.Context) (uint64, uint64) {
	offsetRaw := ctx.DefaultQuery("offset", strconv.FormatInt(DEFAULT_OFFSET, 10))
	limitRaw := ctx.DefaultQuery("limit", strconv.FormatInt(DEFAULT_LIMIT, 10))
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
