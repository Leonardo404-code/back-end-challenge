package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	customErr "frete-rapido-api/pkg/errors"
	"frete-rapido-api/pkg/shipping"
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
	bodyReq := new(shipping.ShippingDataRequest)

	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		customErr.Error(ctx, http.StatusBadRequest, fmt.Errorf("error in parse body: %v", err))
		return
	}

	quotes, err := h.shippingSvc.Quotes(bodyReq)
	if err != nil {
		customErr.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, quotes)
}
