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
	"strconv"
)

type UserController struct {
	userFacade facadeIntf.UserFacade
	logger     loggerIntf.LoggerInterface
}

func NewUserController(userFacade facadeIntf.UserFacade, logger loggerIntf.LoggerInterface) controllerIntf.UserController {
	return &UserController{
		userFacade: userFacade,
		logger:     logger,
	}
}

func (u *UserController) getContextValues(c *gin.Context) context.Context {
	requestID, _ := c.Get("request_id")
	methodRequest, _ := c.Get("method_request")
	urlRequest, _ := c.Get("url_request")

	ctx := context.WithValue(c.Request.Context(), "request_id", requestID.(string))
	ctx = context.WithValue(ctx, "method_request", methodRequest.(string))
	ctx = context.WithValue(ctx, "request_url", urlRequest.(string))

	return ctx
}

// CreateUser creates a new user
// @Summary			Create user
// @Description		Creates a user
// @Tags			User
// @Accept			json
// @Produce			json
// @Param			User		body 		dto.CreateUserRequest	true	"User to be created"
// @Success			201			{object} 	dto.CreateUserResponse
// @Error			400			{object} 	dto.Error
// @Router			/v1/user/	[post]
// @Security		ApiKeyAuth
func (u *UserController) CreateUser(c *gin.Context) {
	ctx := u.getContextValues(c)
	u.logger.LoggerInfo(ctx, "CreateUser", "controller")
	createUserRequestDTO := &dto.CreateUserRequest{}

	err := c.BindJSON(createUserRequestDTO)
	if err != nil {
		u.logger.LoggerError(ctx, constants.ErrorParsingId, "controller")
		c.JSON(http.StatusBadRequest, &dto.Error{Message: constants.ErrorParsingId.Error()})
		return
	}

	createUserResponseDTO, err := u.userFacade.CreateUser(ctx, createUserRequestDTO)
	if err != nil {
		u.logger.LoggerError(ctx, err, "controller")
		c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
		return
	}
	u.logger.LoggerInfo(ctx, "user created", "controller")
	c.JSON(http.StatusCreated, createUserResponseDTO)
}

// GetByIdUser search for a user
// @Summary			Get user
// @Description		Gets a user
// @Tags			User
// @Accept			json
// @Produce			json
// @Param			id		path	string	true	"id"
// @Success			200		{object} dto.GetUserByIDResponse
// @Error			400		{object} dto.Error
// @Error			404		{object} dto.Error
// @Error			500		{object} dto.Error
// @Router			/v1/user/{id}		[get]
// @Security		ApiKeyAuth
func (u *UserController) GetByIdUser(c *gin.Context) {
	ctx := u.getContextValues(c)

	u.logger.LoggerInfo(ctx, "GetByIdUser", "controller")

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		u.logger.LoggerError(ctx, constants.ErrorParsingId, "controller")
		c.JSON(http.StatusBadRequest, &dto.Error{Message: constants.ErrorParsingId.Error()})
		return
	}

	user, err := u.userFacade.GetByIdUser(ctx, id)
	if err != nil {
		if err.Error() == constants.ErrorUserNotFound.Error() {
			u.logger.LoggerError(ctx, constants.ErrorUserNotFound, "controller")
			c.JSON(http.StatusNotFound, &dto.Error{Message: constants.ErrorUserNotFound.Error()})
			return
		}
		u.logger.LoggerError(ctx, err, "controller")
		c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
		return
	}
	u.logger.LoggerInfo(ctx, "user found", "controller")
	c.JSON(http.StatusOK, user)
}

// UpdateUser Updates a user based on id received as part of the URL
// @Summary			Update a user
// @Description 	Updates a user
// @Tags			User
// @Accept			json
// @Produce			json
// @Param           id      path string         true "id"
// @Param			User	body dto.UpdateUserRequest true "User to be updated"
// @Success			204
// @Error			400		{object} dto.Error
// @Error			500		{object} dto.Error
// @Router			/v1/user/{id}	[put]
// @Security		ApiKeyAuth
func (u *UserController) UpdateUser(c *gin.Context) {
	ctx := u.getContextValues(c)

	tokenUserID, found := c.Get("user_id")
	if !found {
		c.JSON(http.StatusInternalServerError, &dto.Error{Message: "don't found token user id key"})
		return
	}

	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: err.Error()})
		return
	}

	updateUserRequestDTO := &dto.UpdateUserRequest{}

	err = c.BindJSON(updateUserRequestDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: constants.ErrorParsingId.Error()})
		return
	}

	if tokenUserID != userID || tokenUserID != updateUserRequestDTO.ID {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "You do not have permission to modify this user."})
		return
	}

	err = u.userFacade.UpdateUser(ctx, userID, updateUserRequestDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
	}

	c.JSON(http.StatusNoContent, nil)
}

// DeleteUser delete a user
// @Summary			Delete a user
// @Description		Delete a user
// @Tags			User
// @Accept			json
// @Produce			json
// @Param 			id		path		string	true	"id"
// @Success			204
// @Error			400		{object} dto.Error
// @Error			500		{object} dto.Error
// @Router			/v1/user/{id}		[delete]
// @Security		ApiKeyAuth
func (u *UserController) DeleteUser(c *gin.Context) {
	ctx := u.getContextValues(c)

	tokenUserID, found := c.Get("user_id")
	if !found {
		c.JSON(http.StatusInternalServerError, &dto.Error{Message: "don't found token user id key"})
		return
	}

	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: constants.ErrorParsingId.Error()})
		return
	}

	if tokenUserID != userID {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "You do not have permission to delete this user."})
		return
	}

	err = u.userFacade.DeleteUser(ctx, userID)
	if err != nil {
		if err.Error() == errors.New("record not found").Error() {
			c.JSON(http.StatusNotFound, &dto.Error{Message: constants.ErrorUserNotFound.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
