package repository

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	repositoryIntf "kickoff/adapter/database/repository/interface"
	"kickoff/internal/model"
)

type UserRepository struct {
	*GenericRepository[model.User]
}

func NewUserRepository(db *gorm.DB) repositoryIntf.UserRepository {
	genericRepo := newGenericRepository[model.User](db)
	return &UserRepository{
		genericRepo,
	}
}

// GetByLogin requests the database to get a user by login.
func (u *UserRepository) GetByLogin(ctx context.Context, user *model.User) (bool, *model.User, error) {
	conn, err := u.GetConnection(ctx)
	if err != nil {
		return false, nil, err
	}

	userReturn := &model.User{}

	if err = conn.Where(model.User{
		Login:          user.Login,
		HashedPassword: user.HashedPassword,
	}).
		First(userReturn).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil, nil
		}
		return false, nil, err
	}

	return true, userReturn, nil
}

func (u *UserRepository) UpdateUser(ctx context.Context, user *model.User) error {
	conn, err := u.GetConnection(ctx)
	if err != nil {
		return err
	}

	updateData := map[string]interface{}{
		"created_at": user.CreatedAt,
		"full_name":  user.FullName,
		"email":      user.Email,
		"login":      user.Login,
	}

	err = conn.Model(&model.User{}).Where("id = ?", user.ID).Updates(updateData).Error
	if err != nil {
		return err
	}
	return nil

}
