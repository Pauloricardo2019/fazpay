package dto

import (
	"kickoff/internal/model"
	"time"
)

type CreateUserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (u *CreateUserRequest) ParseToUserObject() *model.User {
	return &model.User{
		FullName: u.FullName,
		Email:    u.Email,
		Login:    u.Login,
		Password: u.Password,
	}
}

///////////////////////////////////////-------------------//////////////////////////////////////////////////////////

type UpdateUserRequest struct {
	ID       uint64    `json:"id"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	Login    string    `json:"login"`
	CreateAt time.Time `json:"create_at"`
}

func (u *UpdateUserRequest) ParseToUserObject() *model.User {

	return &model.User{
		ID:        u.ID,
		FullName:  u.FullName,
		Email:     u.Email,
		Login:     u.Login,
		CreatedAt: u.CreateAt,
	}
}

///////////////////////////////////////-------------------//////////////////////////////////////////////////////////

type GetUserResponse struct {
	ID       uint64    `json:"id"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	Login    string    `json:"login"`
	CreateAt time.Time `json:"create_at"`
}

func (u *GetUserResponse) ParseFromUserObject(user *model.User) {
	u.ID = user.ID
	u.FullName = user.FullName
	u.Email = user.Email
	u.Login = user.Login
	u.CreateAt = user.CreatedAt
}
