package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	customErr "shipping-calculator-api/pkg/errors"
	"shipping-calculator-api/pkg/shipping"
	servicePkg "shipping-calculator-api/pkg/shipping/service"
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
		handleQuotesErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, quotes)
}

func handleQuotesErr(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, servicePkg.ErrInvalidRequest):
		customErr.Error(ctx, http.StatusBadRequest, err)
		return
	case errors.Is(err, servicePkg.ErrPerformRequest):
		customErr.Error(ctx, http.StatusBadGateway, err)
		return
	default:
		customErr.Error(ctx, http.StatusInternalServerError, err)
		return
	}
}
