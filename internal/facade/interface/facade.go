package facadeIntf

import (
	"context"
	"github.com/Pauloricardo2019/teste_fazpay/internal/dto"
)

type UserFacade interface {
	CreateUser(ctx context.Context, createUserRequestDTO *dto.CreateUserRequest) (*dto.CreateUserResponse, error)
	GetByIdUser(ctx context.Context, id uint64) (*dto.GetUserByIDResponse, error)
	UpdateUser(ctx context.Context, id uint64, updateDTO *dto.UpdateUserRequest) error
	DeleteUser(ctx context.Context, id uint64) error
}

type SecurityFacade interface {
	ValidateToken(ctx context.Context, validateTokenRequest *dto.ValidateTokenRequest) (*dto.ValidateTokenResponse, error)
	Login(ctx context.Context, loginRequest *dto.LoginRequest) (*dto.LoginResponse, error)
}
