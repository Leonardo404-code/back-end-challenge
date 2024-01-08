package handler

import (
	"fmt"
	"net/http"
	"strconv"

	customErr "frete-rapido-api/pkg/errors"
	"frete-rapido-api/pkg/shipping"

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
	lastQuotes := ctx.Query(lastQuotesParam)
	filter := new(shipping.Filter)

	if lastQuotes != "" {
		convertLastQuote, err := strconv.Atoi(lastQuotes)
		if err != nil {
			customErr.Error(ctx, http.StatusBadRequest, fmt.Errorf(`last_quote param must be a number`))
		}

		filter.LastQuotes = convertLastQuote
	}

	metrics, err := h.shippingSvc.Metrics(filter)
	if err != nil {
		customErr.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, metrics)
}
