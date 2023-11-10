package dto

import "kickoff/internal/model"

type ValidateTokenRequest struct {
	Value string `json:"value"`
}

func (v *ValidateTokenRequest) ConvertToVO() *model.Token {
	return &model.Token{
		Value: v.Value,
	}
}

//----------------------------------------------------------

type ValidateTokenResponse struct {
	UserID uint64 `json:"user_id"`
}

func (v *ValidateTokenResponse) ParseFromTokenVO(authorization *model.Token) {
	v.UserID = authorization.UserID
}
