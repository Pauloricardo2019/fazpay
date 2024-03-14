package facade

import (
	"context"
	loggerIntf "github.com/Pauloricardo2019/teste_fazpay/adapter/logger/interface"
	"github.com/Pauloricardo2019/teste_fazpay/internal/constants"
	"github.com/Pauloricardo2019/teste_fazpay/internal/dto"
	facadeIntf "github.com/Pauloricardo2019/teste_fazpay/internal/facade/interface"
	service "github.com/Pauloricardo2019/teste_fazpay/internal/service/interface"
)

type UserFacade struct {
	userService service.UserService
	logger      loggerIntf.LoggerInterface
}

func NewUserFacade(
	userService service.UserService,
	logger loggerIntf.LoggerInterface,
) facadeIntf.UserFacade {
	return &UserFacade{
		userService: userService,
		logger:      logger,
	}
}

func (u *UserFacade) CreateUser(ctx context.Context, createUserRequestDTO *dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	u.logger.LoggerInfo(ctx, "CreateUser", "facade")
	user := createUserRequestDTO.ParseToUserObject()

	found, _, err := u.userService.GetByEmail(ctx, user.Email)
	if err != nil {
		u.logger.LoggerError(ctx, err, "facade")
		return nil, err
	}

	if found {
		u.logger.LoggerWarn(ctx, constants.ErrorEmailAlreadyExists.Error(), "facade")
		return nil, constants.ErrorEmailAlreadyExists
	}

	createdUser, err := u.userService.Create(ctx, user)
	if err != nil {
		u.logger.LoggerError(ctx, err, "facade")
		return nil, err
	}
	u.logger.LoggerInfo(ctx, "user created", "facade")
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

	found, _, err := u.userService.GetByEmail(ctx, updateUser.Email)
	if err != nil {
		return err
	}

	if found {
		return constants.ErrorEmailAlreadyExists
	}

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
