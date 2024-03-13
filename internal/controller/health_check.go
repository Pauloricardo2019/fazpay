package controller

import (
	controllerIntf "github.com/Pauloricardo2019/teste_fazpay/internal/controller/interface"
	"github.com/Pauloricardo2019/teste_fazpay/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthCheckController struct{}

func NewHealthCheckController() controllerIntf.HealthCheckController {
	return &HealthCheckController{}
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
	c.JSON(http.StatusOK, &dto.HealthCheckResponse{
		Status: "OK",
	})
}
