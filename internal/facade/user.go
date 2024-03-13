package facade

import (
	"context"
	"github.com/Pauloricardo2019/teste_fazpay/internal/constants"
	"github.com/Pauloricardo2019/teste_fazpay/internal/dto"
	facadeIntf "github.com/Pauloricardo2019/teste_fazpay/internal/facade/interface"
	service "github.com/Pauloricardo2019/teste_fazpay/internal/service/interface"
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

func (u *UserFacade) CreateUser(ctx context.Context, createUserRequestDTO *dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	user := createUserRequestDTO.ParseToUserObject()

	createdUser, err := u.userService.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	createUserResponse := &dto.CreateUserResponse{}
	createUserResponse.ID = createdUser.ID

	return createUserResponse, nil
}

func (u *UserFacade) GetByIdUser(ctx context.Context, id uint64) (*dto.GetUserByIDResponse, error) {

	found, user, err := u.userService.GetById(ctx, id)
	if err != nil {
		found = false
		return nil, err
	}
	if found != true {
		return nil, constants.ErrorUserNotFound
	}

	getByIdUserResponse := &dto.GetUserByIDResponse{}
	getByIdUserResponse.ParseFromUserObject(user)

	return getByIdUserResponse, nil
}

func (u *UserFacade) UpdateUser(ctx context.Context, id uint64, updateUserRequest *dto.UpdateUserRequest) error {

	updateUserRequest.ID = id
	updateUser := updateUserRequest.ParseToUserObject()

	if err := u.userService.Update(ctx, updateUser); err != nil {
		return err
	}
	return nil
}

func (u *UserFacade) DeleteUser(ctx context.Context, id uint64) error {

	if err := u.userService.Delete(ctx, id); err != nil {
		return err
	}
	return nil

}
