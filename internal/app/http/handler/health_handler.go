package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status string `json:"status"`
}

// Healthcheck
// @Summary      Healthcheck
// @Description  Check if the API is online
// @Tags         system
// @Produce      json
// @Success      200  {object}  HealthResponse
// @Router       /health [get]
func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, HealthResponse{Status: "ok"})
}
