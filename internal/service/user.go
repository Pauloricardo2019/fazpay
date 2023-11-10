package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"kickoff/adapter/database/repository/interface"
	"kickoff/internal/model"
	serviceIntf "kickoff/internal/service/interface"
)

type UserService struct {
	userRepository repositoryIntf.UserRepository
}

func NewUserService(userRepository repositoryIntf.UserRepository) serviceIntf.UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

// Create requests the repository layer to create a new user.
func (u *UserService) Create(ctx context.Context, user *model.User) (*model.User, error) {
	user = encryptAndClearingPassword(user)
	return u.userRepository.Create(ctx, user)
}

// GetList requests the repository layer to get all users.
func (u *UserService) GetList(ctx context.Context, limit int, offset int) ([]model.User, error) {
	return u.userRepository.GetList(ctx, limit, offset)
}

// GetCount requests the repository layer to get all users.
func (u *UserService) GetCount(ctx context.Context) (int64, error) {
	return u.userRepository.GetCount(ctx)
}

// GetById requests the repository layer to get a user.
func (u *UserService) GetById(ctx context.Context, id uint64) (bool, *model.User, error) {
	return u.userRepository.GetById(ctx, id)
}

// Update requests the repository layer to update a user.
func (u *UserService) Update(ctx context.Context, user *model.User) error {
	return u.userRepository.Update(ctx, user)
}

// Delete requests the repository layer to delete a user.
func (u *UserService) Delete(ctx context.Context, id uint64) error {
	return u.userRepository.Delete(ctx, id)
}

// GetByLogin - requests the repository layer to get a user.
func (u *UserService) GetByLogin(ctx context.Context, user *model.User) (bool, *model.User, error) {
	user = encryptAndClearingPassword(user)
	return u.userRepository.GetByLogin(ctx, user)
}

func encryptAndClearingPassword(usuario *model.User) *model.User {
	sum := sha256.Sum256([]byte(usuario.Password))
	usuario.HashedPassword = fmt.Sprintf("%x", sum)

	return usuario
}
