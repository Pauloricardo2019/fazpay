package repository

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	repositoryIntf "kickoff/adapter/database/repository/interface"
	"kickoff/internal/model"
)

type TokenRepository struct {
	*GenericRepository[model.Token]
}

func NewTokenRepository(db *gorm.DB) repositoryIntf.TokenRepository {
	genericRepo := newGenericRepository[model.Token](db)
	return &TokenRepository{
		genericRepo,
	}
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
