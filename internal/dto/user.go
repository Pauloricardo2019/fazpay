package dto

import (
	"github.com/Pauloricardo2019/teste_fazpay/internal/model"
	"time"
)

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
} //@name CreateUserRequest

type CreateUserResponse struct {
	ID uint64 `json:"id"`
} //@name CreateUserResponse

func (c *CreateUserRequest) ParseToUserObject() *model.User {
	return &model.User{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		Password:  c.Password,
	}
}

func (u *CreateUserRequest) ParseToUserVO() *model.User {
	return &model.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
	}
}

func (u *CreateUserResponse) ParseFromUserVO(user *model.User) {
	u.ID = user.ID
}

type UpdateUserRequest struct {
	ID        uint64    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} //@name UpdateUserRequest

func (u *UpdateUserRequest) ParseToUserObject() *model.User {
	return &model.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *UpdateUserRequest) ParseToUpdateUserVO() *model.User {

	return &model.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
	}
}

type GetUserByIDResponse struct {
	ID        uint64    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Created   time.Time `json:"created_at"`
	Updated   time.Time `json:"updated_at"`
} //@name GetUserByIDResponse

func (u *GetUserByIDResponse) ParseFromUserObject(user *model.User) {
	u.ID = user.ID
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email
	u.Created = user.CreatedAt
	u.Updated = user.UpdatedAt
}
