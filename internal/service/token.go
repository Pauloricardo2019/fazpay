package service

import (
	"context"
	"github.com/gofrs/uuid"
	"kickoff/adapter/database/repository/interface"
	"kickoff/internal/constants"
	"kickoff/internal/model"
	serviceIntf "kickoff/internal/service/interface"
)

type TokenService struct {
	tokenRepository repositoryIntf.TokenRepository
}

func NewTokenService(tokenRepository repositoryIntf.TokenRepository) serviceIntf.TokenService {
	return &TokenService{
		tokenRepository: tokenRepository,
	}
}

func (u *TokenService) CreateForLogin(ctx context.Context, user *model.User) (*model.Token, error) {
	tokenUUID, err := uuid.NewV4()
	if err != nil {
		return nil, constants.ErrorCreateToken
	}

	token := &model.Token{
		Value:  tokenUUID.String(),
		UserID: user.ID,
	}

	return u.tokenRepository.Create(ctx, token)
}

func (u *TokenService) Create(ctx context.Context, token *model.Token) (*model.Token, error) {
	return u.tokenRepository.Create(ctx, token)
}

func (u *TokenService) GetByValue(ctx context.Context, value string) (bool, *model.Token, error) {
	return u.tokenRepository.GetByValue(ctx, value)
}
