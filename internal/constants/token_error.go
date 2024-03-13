package constants

import "errors"

var (
	ErrorTokenExpired    = errors.New("token expired")
	ErrorTokenValueEmpty = errors.New("token value not filled")
	ErrorTokenNotFound   = errors.New("token not found")
)
