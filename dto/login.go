package dto

import "kickoff/internal/model"

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (l *LoginRequest) ConvertToVO() *model.User {
	return &model.User{
		Login:    l.Login,
		Password: l.Password,
	}
}

// -------------

type LoginResponse struct {
	Token    string `json:"token"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	UserID   uint64 `json:"user_id"`
}

func (t *LoginResponse) ParseFromTokenAndUserVO(token *model.Token, user *model.User) {
	t.Token = token.Value
	t.FullName = user.FullName
	t.Email = user.Email
	t.UserID = token.UserID

}
