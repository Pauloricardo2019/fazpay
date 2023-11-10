package facade

import (
	"errors"
	"golang.org/x/net/context"
	"kickoff/dto"
	"kickoff/internal/constants"
	facadeIntf "kickoff/internal/facade/interface"
	service "kickoff/internal/service/interface"
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
		return nil, errors.New("token value not filled")
	}

	found, tokenFound, err := s.tokenService.GetByValue(ctx, tokenVO.Value)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("token not found")
	}

	tokenResponse := &dto.ValidateTokenResponse{}
	tokenResponse.ParseFromTokenVO(tokenFound)

	return tokenResponse, nil
}

func (s *SecurityFacade) Login(ctx context.Context, loginRequest *dto.LoginRequest) (*dto.LoginResponse, error) {
	loginUserObject := loginRequest.ConvertToVO()

	if loginRequest.Login == "" {
		return nil, errors.New("login not filled")
	}

	if loginRequest.Password == "" {
		return nil, errors.New("password not filled")
	}

	found, userFound, err := s.userService.GetByLogin(ctx, loginUserObject)
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
