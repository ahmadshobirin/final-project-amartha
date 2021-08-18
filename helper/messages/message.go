package messages

import "errors"

var (
	// error_messages
	ErrIDNotFound               = errors.New("id not found")
	ErrDuplicateData            = errors.New("duplicate data")
	ErrDataAlreadyExist         = errors.New("data already exist")
	ErrInternalServer           = errors.New("something gone wrong, contact administrator")
	ErrUsernamePasswordNotFound = errors.New("(Username) or (Password) empty")
	ErrNotFound                 = errors.New("data not found")
	ErrInvalidBearerToken       = errors.New("invalid_bearer_token")
	ErrExpiredToken             = errors.New("expired_token")
	ErrInvalidRole              = errors.New("invalid_role")
	ErrInvalidCred              = errors.New("invalid_credential")
	//Modular
	BaseResponseMessageSuccess = "success"
	BaseResponseMessageFailed  = "something not right"
)
