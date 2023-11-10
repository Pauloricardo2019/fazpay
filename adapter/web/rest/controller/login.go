package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	controllerIntf "kickoff/adapter/web/rest/controller/interface"
	"kickoff/dto"
	"kickoff/internal/constants"
	facadeIntf "kickoff/internal/facade/interface"
	"net/http"
)

type loginController struct {
	securityFacade facadeIntf.SecurityFacade
}

func NewLoginController(securityFacade facadeIntf.SecurityFacade) controllerIntf.LoginController {
	return &loginController{
		securityFacade: securityFacade,
	}
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
	ctx := c.Request.Context()

	loginRequest := &dto.LoginRequest{}

	if err := c.BindJSON(loginRequest); err != nil {
		c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
		return
	}

	loginResponse, err := l.securityFacade.Login(ctx, loginRequest)

	if err != nil {
		switch {
		case errors.Is(err, constants.ErrorUserNotFound):
			c.JSON(http.StatusNotFound, &dto.Error{Message: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, loginResponse)

}
