package facade

import (
	"context"
	"kickoff/dto"
	"kickoff/internal/constants"
	facadeIntf "kickoff/internal/facade/interface"
	service "kickoff/internal/service/interface"
)

type UserFacade struct {
	userService service.UserService
}

func NewUserFacade(
	userService service.UserService,
) facadeIntf.UserFacade {
	return &UserFacade{
		userService: userService,
	}
}

// CreateUser requests the service layer to create a new user.
func (u *UserFacade) CreateUser(ctx context.Context, createUserRequestDTO *dto.CreateUserRequest) (*dto.CreateResponse, error) {
	user := createUserRequestDTO.ParseToUserObject()

	createdUser, err := u.userService.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	createUserResponse := &dto.CreateResponse{}
	createUserResponse.ID = createdUser.ID

	return createUserResponse, nil
}

// GetByIdUser requests the service layer to get a user.
func (u *UserFacade) GetByIdUser(ctx context.Context, id uint64) (*dto.GetUserResponse, error) {

	found, user, err := u.userService.GetById(ctx, id)
	if err != nil {
		found = false
		return nil, err
	}
	if found != true {
		return nil, constants.ErrorUserNotFound
	}

	getByIdUserResponse := &dto.GetUserResponse{}
	getByIdUserResponse.ParseFromUserObject(user)

	return getByIdUserResponse, nil
}

// UpdateUser requests the service layer to update a user.
func (u *UserFacade) UpdateUser(ctx context.Context, id uint64, updateUserRequest *dto.UpdateUserRequest) error {

	updateUserRequest.ID = id
	updateUser := updateUserRequest.ParseToUserObject()

	if err := u.userService.Update(ctx, updateUser); err != nil {
		return err
	}
	return nil
}

// DeleteUser requests the service layer to delete a user.
func (u *UserFacade) DeleteUser(ctx context.Context, id uint64) error {

	if err := u.userService.Delete(ctx, id); err != nil {
		return err
	}
	return nil

}
