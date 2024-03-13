package serviceIntf

import (
	"context"
	"github.com/Pauloricardo2019/teste_fazpay/internal/model"
)

type SaveBaseService interface {
	BeginTransaction(ctx context.Context) (context.Context, error)
	RollbackTransaction(ctx context.Context) error
	CommitTransaction(ctx context.Context) error
}

type TransactionHandler func(ctx context.Context) error

type TransactionService interface {
	DoWork(ctx context.Context, transactionHandler TransactionHandler) error
}

type UserService interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	GetById(ctx context.Context, id uint64) (bool, *model.User, error)
	GetByEmailAndPassword(ctx context.Context, user *model.User) (bool, *model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint64) error
	GetByEmail(ctx context.Context, email string) (bool, *model.User, error)
}

type TokenService interface {
	CreateForLogin(ctx context.Context, user *model.User) (*model.Token, error)
	GetByValue(ctx context.Context, value string) (bool, *model.Token, error)
}
