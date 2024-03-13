package dto

import "github.com/Pauloricardo2019/teste_fazpay/internal/model"

type ValidateTokenRequest struct {
	Value string `json:"value"`
} //@name ValidateTokenRequest

func (v *ValidateTokenRequest) ConvertToVO() *model.Token {
	return &model.Token{
		Value: v.Value,
	}
}

type ValidateTokenResponse struct {
	UserID uint64 `json:"user_id"`
} //@name ValidateTokenResponse

func (v *ValidateTokenResponse) ParseFromTokenVO(authorization *model.Token) {
	v.UserID = authorization.UserID
}
