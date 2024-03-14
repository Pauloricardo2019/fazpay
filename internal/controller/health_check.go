package controller

import (
	"context"
	loggerIntf "github.com/Pauloricardo2019/teste_fazpay/adapter/logger/interface"
	controllerIntf "github.com/Pauloricardo2019/teste_fazpay/internal/controller/interface"
	"github.com/Pauloricardo2019/teste_fazpay/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthCheckController struct {
	logger loggerIntf.LoggerInterface
}

func NewHealthCheckController(logger loggerIntf.LoggerInterface) controllerIntf.HealthCheckController {
	return &HealthCheckController{
		logger: logger,
	}
}

func (h *HealthCheckController) getContextValues(c *gin.Context) context.Context {
	requestID, _ := c.Get("request_id")
	methodRequest, _ := c.Get("method_request")
	urlRequest, _ := c.Get("url_request")

	ctx := context.WithValue(c.Request.Context(), "request_id", requestID.(string))
	ctx = context.WithValue(ctx, "method_request", methodRequest.(string))
	ctx = context.WithValue(ctx, "request_url", urlRequest.(string))

	return ctx
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
	ctx := h.getContextValues(c)
	h.logger.LoggerInfo(ctx, "Health-Check", "controller")
	c.JSON(http.StatusOK, &dto.HealthCheckResponse{
		Status: "OK",
	})
}
