package service

import (
	"context"
	loggerIntf "github.com/Pauloricardo2019/teste_fazpay/adapter/logger/interface"
	"github.com/Pauloricardo2019/teste_fazpay/internal/constants"
	"github.com/Pauloricardo2019/teste_fazpay/internal/model"
	"github.com/Pauloricardo2019/teste_fazpay/internal/repository/interface"
	serviceIntf "github.com/Pauloricardo2019/teste_fazpay/internal/service/interface"
	"github.com/gofrs/uuid"
	"time"
)

type TokenService struct {
	tokenRepository repositoryIntf.TokenRepository
	logger          loggerIntf.LoggerInterface
}

func NewTokenService(tokenRepository repositoryIntf.TokenRepository, logger loggerIntf.LoggerInterface) serviceIntf.TokenService {
	return &TokenService{
		tokenRepository: tokenRepository,
		logger:          logger,
	}
}

func (u *TokenService) CreateForLogin(ctx context.Context, user *model.User) (*model.Token, error) {
	u.logger.LoggerInfo(ctx, "CreateForLogin", "service")
	tokenUUID, err := uuid.NewV4()
	if err != nil {
		return nil, constants.ErrorCreateToken
	}

	token := &model.Token{
		Value:     tokenUUID.String(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}
	u.logger.LoggerInfo(ctx, "create token object", "service")
	return u.tokenRepository.Create(ctx, token)
}

func (u *TokenService) Create(ctx context.Context, token *model.Token) (*model.Token, error) {
	u.logger.LoggerInfo(ctx, "Create", "service")
	return u.tokenRepository.Create(ctx, token)
}

func (u *TokenService) GetByValue(ctx context.Context, value string) (bool, *model.Token, error) {
	return u.tokenRepository.GetByValue(ctx, value)
}
