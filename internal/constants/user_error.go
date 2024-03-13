package constants

import "errors"

var (
	ErrorUserNotFound        = errors.New("user not found")
	PassCantBeEmpty          = errors.New("password can't be empty")
	PassMin7Characters       = errors.New("password must have at least 7 characters")
	PassMin1Letter           = errors.New("password must have at least 1 letter")
	PassMin1LetterUpper      = errors.New("password must have at least 1 letter uppercase")
	PassMin1Number           = errors.New("password must have at least 1 number")
	PassMin1SpecialCharacter = errors.New("password must have at least 1 special character")
	InvalidCPF               = errors.New("invalid cpf")
	LoginInvalid             = errors.New("incorrect cpf or password")
	CPFAlreadyExists         = errors.New("this cpf has already been registered")
	InvalidEmail             = errors.New("invalid email")
	EmailAlreadyUsed         = errors.New("this email has already been used")
	ErrorParsingId           = errors.New("error: error parsing id")
	ErrorCreateToken         = errors.New("error: cannot create a token")
)
