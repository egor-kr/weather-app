package weather_transport_http

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) GetForecast(c *gin.Context) {
	city := c.Param("city")

	forecast, err := h.service.GetForecast(c.Request.Context(), city)
	if err != nil {
		slog.Error("get forecast:", slog.String("err", err.Error()))
		c.JSON(500, gin.H{"error": "Failed to get forecast"})
		return
	}

	resp := domainToResponse(forecast)
	c.JSON(200, resp)
}
