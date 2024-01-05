package handler

import (
	"github.com/gin-gonic/gin"
)

// @Summary Get quotes data in database
// @Description Endpoint to search quotes data in database
// @Tags shipping
// @Param last_quotes query string false "Last Quotes"
// @Router /metrics [get]
// @Produce json
// @Success 200 {object} handler.QuotesResponseDoc
func (h *handler) Metrics(ctx *gin.Context) {
	panic("not implemented")
}
