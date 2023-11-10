package repositoryIntf

import (
	"context"
	"kickoff/internal/model"
)

type DBRepository interface {
	BeginTransaction(ctx context.Context) (context.Context, error)
	RollbackTransaction(ctx context.Context) error
	CommitTransaction(ctx context.Context) error
}

type TokenRepository interface {
	Create(ctx context.Context, token *model.Token) (*model.Token, error)
	GetByValue(ctx context.Context, value string) (bool, *model.Token, error)
}

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	GetCount(ctx context.Context) (int64, error)
	GetList(ctx context.Context, limit int, offset int) ([]model.User, error)
	GetById(ctx context.Context, id uint64) (bool, *model.User, error)
	GetByLogin(ctx context.Context, user *model.User) (bool, *model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint64) error
}

type Migrator interface {
	ExecuteMigrations(ctx context.Context) error
}
