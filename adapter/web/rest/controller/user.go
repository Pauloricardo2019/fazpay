package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	controllerIntf "kickoff/adapter/web/rest/controller/interface"
	"kickoff/dto"
	"kickoff/internal/constants"
	facadeIntf "kickoff/internal/facade/interface"
	"net/http"
	"strconv"
)

type UserController struct {
	userFacade facadeIntf.UserFacade
}

func NewUserController(userFacade facadeIntf.UserFacade) controllerIntf.UserController {
	return &UserController{
		userFacade: userFacade,
	}
}

// CreateUser creates a new user
// @Summary			Create user
// @Description		Creates a user
// @Tags			User
// @Accept			json
// @Produce			json
// @Param			User		body 		dto.CreateUserRequest	true	"User to be created"
// @Success			201			{object} 	dto.CreateResponse
// @Error			400			{object} 	dto.Error
// @Router			/v1/user/	[post]
// @Security		ApiKeyAuth
func (u *UserController) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	createUserRequestDTO := &dto.CreateUserRequest{}

	err := c.BindJSON(createUserRequestDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: constants.ErrorParsingId.Error()})
		return
	}

	createUserResponseDTO, err := u.userFacade.CreateUser(ctx, createUserRequestDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createUserResponseDTO)
}

// GetByIdUser search for a user
// @Summary			Get user
// @Description		Gets a user
// @Tags			User
// @Accept			json
// @Produce			json
// @Param			id		path	string	true	"id"
// @Success			200		{object} dto.GetUserResponse
// @Error			400		{object} dto.Error
// @Error			404		{object} dto.Error
// @Error			500		{object} dto.Error
// @Router			/v1/user/{id}		[get]
// @Security		ApiKeyAuth
func (u *UserController) GetByIdUser(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: constants.ErrorParsingId.Error()})
		return
	}

	user, err := u.userFacade.GetByIdUser(ctx, id)
	if err != nil {
		if err.Error() == constants.ErrorUserNotFound.Error() {
			c.JSON(http.StatusNotFound, &dto.Error{Message: constants.ErrorUserNotFound.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
		return
	}

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
	ctx := c.Request.Context()

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
	ctx := c.Request.Context()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: constants.ErrorParsingId.Error()})
		return
	}

	err = u.userFacade.DeleteUser(ctx, id)
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
