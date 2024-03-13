package service

import (
	"context"
	"github.com/Pauloricardo2019/teste_fazpay/internal/constants"
	"github.com/Pauloricardo2019/teste_fazpay/internal/model"
	"github.com/Pauloricardo2019/teste_fazpay/internal/repository/interface"
	serviceIntf "github.com/Pauloricardo2019/teste_fazpay/internal/service/interface"
	"github.com/gofrs/uuid"
	"time"
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
		Value:     tokenUUID.String(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	return u.tokenRepository.Create(ctx, token)
}

func (u *TokenService) Create(ctx context.Context, token *model.Token) (*model.Token, error) {
	return u.tokenRepository.Create(ctx, token)
}

func (u *TokenService) GetByValue(ctx context.Context, value string) (bool, *model.Token, error) {
	return u.tokenRepository.GetByValue(ctx, value)
}
