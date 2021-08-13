package messages

import "errors"

var (
	// error_messages
	ErrIDNotFound       = errors.New("id not found")
	ErrDuplicateData    = errors.New("duplicate data")
	ErrDataAlreadyExist = errors.New("data already exist")
	ErrInternalServer   = errors.New("something gone wrong, contact administrator")

	//Modular
	BaseResponseMessageSuccess = "success"
	BaseResponseMessageFailed  = "something not right"
)
