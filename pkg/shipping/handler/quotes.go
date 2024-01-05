package handler

import (
	"github.com/gin-gonic/gin"
)

// @Summary Fetch quote data and persist in database
// @Description Endpoint to fetch quotes data from the Frete rápido API and persists it in the database
// @Tags shipping
// @Param shippingData body handler.ShippingDataDoc true "body request"
// @Router /quotes [post]
// @Accept json
// @Produce json
// @Success 200 {object} handler.QuotesResponseDoc
func (h *handler) Quotes(ctx *gin.Context) {
	panic("not implemented")
}
