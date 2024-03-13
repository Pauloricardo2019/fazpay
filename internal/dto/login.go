package dto

import "github.com/Pauloricardo2019/teste_fazpay/internal/model"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
} //@name LoginRequest

func (l *LoginRequest) ConvertToVO() *model.User {
	return &model.User{
		Email:    l.Email,
		Password: l.Password,
	}
}

type LoginResponse struct {
	Token     string `json:"token"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	UserID    uint64 `json:"user_id"`
} //@name LoginResponse

func (t *LoginResponse) ParseFromTokenAndUserVO(token *model.Token, user *model.User) {
	t.Token = token.Value
	t.FirstName = user.FirstName
	t.LastName = user.LastName
	t.Email = user.Email
	t.UserID = token.UserID

}
