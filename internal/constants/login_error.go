package constants

import "errors"

var (
	ErrorLoginValueEmpty = errors.New("login value not filled")
	ErrorLoginPassEmpty  = errors.New("password value not filled")
)
