package facade

import (
	"github.com/Pauloricardo2019/teste_fazpay/internal/constants"
	"github.com/Pauloricardo2019/teste_fazpay/internal/dto"
	facadeIntf "github.com/Pauloricardo2019/teste_fazpay/internal/facade/interface"
	service "github.com/Pauloricardo2019/teste_fazpay/internal/service/interface"
	"golang.org/x/net/context"
	"time"
)

type SecurityFacade struct {
	tokenService service.TokenService
	userService  service.UserService
}

func NewSecurityFacade(
	tokenService service.TokenService,
	userService service.UserService,
) facadeIntf.SecurityFacade {
	return &SecurityFacade{
		tokenService: tokenService,
		userService:  userService,
	}
}

func (s *SecurityFacade) ValidateToken(ctx context.Context, validateTokenRequest *dto.ValidateTokenRequest) (*dto.ValidateTokenResponse, error) {
	tokenVO := validateTokenRequest.ConvertToVO()

	if tokenVO.Value == "" {
		return nil, constants.ErrorTokenValueEmpty
	}

	found, tokenFound, err := s.tokenService.GetByValue(ctx, tokenVO.Value)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, constants.ErrorTokenNotFound
	}

	expiresAtIsValid := s.validateExpiresAt(tokenFound.ExpiresAt)
	if !expiresAtIsValid {
		return nil, constants.ErrorTokenExpired
	}

	tokenResponse := &dto.ValidateTokenResponse{}
	tokenResponse.ParseFromTokenVO(tokenFound)

	return tokenResponse, nil
}

func (s *SecurityFacade) validateExpiresAt(expiresAt time.Time) bool {
	if time.Now().Sub(expiresAt) > time.Hour*24 {
		return false
	}
	return true
}

func (s *SecurityFacade) Login(ctx context.Context, loginRequest *dto.LoginRequest) (*dto.LoginResponse, error) {
	loginUserObject := loginRequest.ConvertToVO()

	if loginRequest.Email == "" {
		return nil, constants.ErrorLoginValueEmpty
	}

	if loginRequest.Password == "" {
		return nil, constants.ErrorLoginPassEmpty
	}

	found, userFound, err := s.userService.GetByEmailAndPassword(ctx, loginUserObject)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, constants.ErrorUserNotFound
	}

	tokenCreated, err := s.tokenService.CreateForLogin(ctx, userFound)
	if err != nil {
		return nil, err
	}

	loginResponse := &dto.LoginResponse{}
	loginResponse.ParseFromTokenAndUserVO(tokenCreated, userFound)

	return loginResponse, nil
}
