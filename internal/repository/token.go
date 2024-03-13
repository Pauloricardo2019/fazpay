package repository

import (
	"context"
	"github.com/Pauloricardo2019/teste_fazpay/internal/model"
	repositoryIntf "github.com/Pauloricardo2019/teste_fazpay/internal/repository/interface"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type TokenRepository struct {
	*BaseRepository
}

func NewTokenRepository(db *gorm.DB) repositoryIntf.TokenRepository {
	baseRepo := NewBaseRepository(db)
	return &TokenRepository{
		baseRepo,
	}
}

func (u *TokenRepository) Create(ctx context.Context, token *model.Token) (*model.Token, error) {
	conn, err := u.GetConnection(ctx)
	if err != nil {
		return nil, err
	}

	if err = conn.Create(token).Error; err != nil {
		return nil, err
	}

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
