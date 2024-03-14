package controller

import (
	"context"
	"errors"
	loggerIntf "github.com/Pauloricardo2019/teste_fazpay/adapter/logger/interface"
	"github.com/Pauloricardo2019/teste_fazpay/internal/constants"
	controllerIntf "github.com/Pauloricardo2019/teste_fazpay/internal/controller/interface"
	"github.com/Pauloricardo2019/teste_fazpay/internal/dto"
	facadeIntf "github.com/Pauloricardo2019/teste_fazpay/internal/facade/interface"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginController struct {
	securityFacade facadeIntf.SecurityFacade
	logger         loggerIntf.LoggerInterface
}

func NewLoginController(securityFacade facadeIntf.SecurityFacade, logger loggerIntf.LoggerInterface) controllerIntf.LoginController {
	return &loginController{
		securityFacade: securityFacade,
		logger:         logger,
	}
}

func (l *loginController) getContextValues(c *gin.Context) context.Context {
	requestID, _ := c.Get("request_id")
	methodRequest, _ := c.Get("method_request")
	urlRequest, _ := c.Get("url_request")

	ctx := context.WithValue(c.Request.Context(), "request_id", requestID.(string))
	ctx = context.WithValue(ctx, "method_request", methodRequest.(string))
	ctx = context.WithValue(ctx, "request_url", urlRequest.(string))

	return ctx
}

// Login - Perform user login
// @Summary - login user
// @Description - Performs user login and returns a token
// @Tags Login
// @Accept json
// @Produce json
// @Param loginRequest body dto.LoginRequest true "login to be performed"
// @Success 200 {object} dto.LoginResponse
// @Failure 404 {object} dto.Error
// @Failure 500 {object} dto.Error
// @Router /v1/auth/login/ [post]
// @Security ApiKeyAuth
func (l *loginController) Login(c *gin.Context) {
	ctx := l.getContextValues(c)
	l.logger.LoggerInfo(ctx, "Login", "controller")
	loginRequest := &dto.LoginRequest{}

	if err := c.BindJSON(loginRequest); err != nil {
		l.logger.LoggerError(ctx, err, "controller")
		c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
		return
	}

	loginResponse, err := l.securityFacade.Login(ctx, loginRequest)
	if err != nil {
		switch {
		case errors.Is(err, constants.ErrorUserNotFound):
			l.logger.LoggerError(ctx, constants.ErrorUserNotFound, "controller")
			c.JSON(http.StatusNotFound, &dto.Error{Message: err.Error()})
		default:
			l.logger.LoggerError(ctx, err, "controller")
			c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
		}
		return
	}

	l.logger.LoggerInfo(ctx, "user logged in", "controller")
	c.JSON(http.StatusOK, loginResponse)

}
