package repository

import (
	"context"
	loggerIntf "github.com/Pauloricardo2019/teste_fazpay/adapter/logger/interface"
	"github.com/Pauloricardo2019/teste_fazpay/internal/model"
	repositoryIntf "github.com/Pauloricardo2019/teste_fazpay/internal/repository/interface"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type TokenRepository struct {
	*BaseRepository
	logger loggerIntf.LoggerInterface
}

func NewTokenRepository(db *gorm.DB, logger loggerIntf.LoggerInterface) repositoryIntf.TokenRepository {
	baseRepo := NewBaseRepository(db)
	return &TokenRepository{
		baseRepo,
		logger,
	}
}

func (u *TokenRepository) Create(ctx context.Context, token *model.Token) (*model.Token, error) {
	u.logger.LoggerInfo(ctx, "Create", "repository")
	conn, err := u.GetConnection(ctx)
	if err != nil {
		return nil, err
	}

	if err = conn.Create(token).Error; err != nil {
		return nil, err
	}
	u.logger.LoggerInfo(ctx, "token created", "repository")
	return token, nil
}

func (u *TokenRepository) GetByValue(ctx context.Context, value string) (bool, *model.Token, error) {
	conn, err := u.GetConnection(ctx)
	if err != nil {
		return false, nil, err
	}

	token := &model.Token{}

	if err = conn.Where(model.Token{
		Value: value,
	}).First(token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, token, nil
		}
		return false, nil, err
	}

	return true, token, nil
}
