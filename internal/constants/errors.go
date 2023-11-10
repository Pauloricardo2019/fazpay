package constants

import "errors"

var (
	ErrorPersonalInfoNotFound  = errors.New("error: personal info not found")
	ErrorAddressNotFound       = errors.New("error: address not found")
	ErrorGymNotFound           = errors.New("error: gym not found")
	ErrorBusinessInfoNotFound  = errors.New("error: business info not found")
	ErrorClientNotFound        = errors.New("error: client not found")
	ErrorTokenNotFound         = errors.New("error: token not found")
	ErrorUserNotFound          = errors.New("error: user not found")
	ErrorPhoneNotFound         = errors.New("error: phone not found")
	ErrorModalityNotFound      = errors.New("error: modality not found")
	ErrorParsingJson           = errors.New("error: error parsing json")
	ErrorParsingId             = errors.New("error: error parsing id")
	ErrorCreateToken           = errors.New("error: cannot create a token")
	ErrorEmptyToken            = errors.New("error: token cannot be empty")
	ErrorUserLoginNotFilled    = errors.New("error: login required")
	ErrorUserPasswordNotFilled = errors.New("error: password required")
	ErrorCustomerNotFound      = errors.New("error: customer not found")
	ErrorPhoneActionNotValid   = errors.New("error: phone action invalid")
	ErrorAddressActionNotValid = errors.New("error: address action invalid")
)
