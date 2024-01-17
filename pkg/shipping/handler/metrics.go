package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	customErr "shipping-calculator-api/pkg/errors"
	"shipping-calculator-api/pkg/shipping"
	servicePkg "shipping-calculator-api/pkg/shipping/service"

	"github.com/gin-gonic/gin"
)

// @Summary Get quotes data in database
// @Description Endpoint to search quotes data in database
// @Tags shipping
// @Param last_quotes query int false "Informs the number of quotes that will be returned"
// @Router /metrics [get]
// @Produce json
// @Success 200 {object} handler.MetricsResposneDoc
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
		handlerMetricsErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, metrics)
}

func handlerMetricsErr(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, servicePkg.ErrNotFound):
		customErr.Error(ctx, http.StatusNotFound, err)
		return
	default:
		customErr.Error(ctx, http.StatusInternalServerError, err)
		return
	}
}
