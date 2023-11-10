package controller

import (
	"github.com/gin-gonic/gin"
	logIntf "kickoff/adapter/log/interface"
	controllerIntf "kickoff/adapter/web/rest/controller/interface"
	"kickoff/dto"
	"net/http"
)

type HealthCheckController struct {
	logger logIntf.Logger
}

func NewHealthCheckController(logger logIntf.Logger) controllerIntf.HealthCheckController {
	return &HealthCheckController{
		logger: logger,
	}
}

// HealthCheck - health-check for the server
// @Summary - Health-Check
// @Description - Health-Check for the API
// @Tags Health-Check
// @Accept json
// @Produce json
// @Success 200 {object} dto.HealthCheckResponse
// @Router /health [get]
func (h *HealthCheckController) HealthCheck(c *gin.Context) {
	ctx := c.Request.Context()

	h.logger.Info(ctx, "Starting")

	c.JSON(http.StatusOK, &dto.HealthCheckResponse{
		Status: "OK",
	})
}
