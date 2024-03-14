package repository

import (
	"context"
	"errors"
	loggerIntf "github.com/Pauloricardo2019/teste_fazpay/adapter/logger/interface"
	"github.com/Pauloricardo2019/teste_fazpay/internal/model"
	repositoryIntf "github.com/Pauloricardo2019/teste_fazpay/internal/repository/interface"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	*BaseRepository
	logger loggerIntf.LoggerInterface
}

func NewUserRepository(db *gorm.DB, logger loggerIntf.LoggerInterface) repositoryIntf.UserRepository {
	baseRepo := NewBaseRepository(db)
	return &UserRepository{
		baseRepo,
		logger,
	}
}

func (u *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	u.logger.LoggerInfo(ctx, "Create", "repository")
	conn, err := u.GetConnection(ctx)
	if err != nil {
		return nil, err
	}

	if err = conn.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetById(ctx context.Context, id uint64) (bool, *model.User, error) {
	conn, err := u.GetConnection(ctx)
	if err != nil {
		return false, nil, err
	}

	user := model.User{}

	if err = conn.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil, nil
		}
		return false, nil, err
	}

	return true, &user, nil
}

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (bool, *model.User, error) {
	u.logger.LoggerInfo(ctx, "GetByEmail", "repository")
	conn, err := u.GetConnection(ctx)
	if err != nil {
		return false, nil, err
	}

	user := model.User{}

	if err = conn.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil, nil
		}
		return false, nil, err
	}

	return true, &user, nil
}

func (u *UserRepository) Update(ctx context.Context, user *model.User) error {
	conn, err := u.GetConnection(ctx)
	if err != nil {
		return err
	}

	updateData := map[string]interface{}{
		"created_at":      user.CreatedAt,
		"updated_at":      time.Now(),
		"first_name":      user.FirstName,
		"last_name":       user.LastName,
		"email":           user.Email,
		"hashed_password": user.HashedPassword,
	}

	err = conn.Model(&model.User{}).Where("id = ?", user.ID).Updates(updateData).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) Delete(ctx context.Context, id uint64) error {
	conn, err := u.GetConnection(ctx)
	if err != nil {
		return err
	}

	if err = conn.Delete(&model.User{ID: id}).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetByEmailAndPassword(ctx context.Context, user *model.User) (bool, *model.User, error) {
	conn, err := u.GetConnection(ctx)
	if err != nil {
		return false, nil, err
	}

	if err = conn.
		Where(&model.User{
			Email:          user.Email,
			HashedPassword: user.HashedPassword,
		}).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil, nil
		}
		return false, nil, err
	}

	return true, user, nil
}
