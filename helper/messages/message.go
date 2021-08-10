package messages

import "errors"

var (
	// error_messages
	ErrIDNotFound        = errors.New("id not found")
	ErrNewsIDResource    = errors.New("(NewsID) not found or empty")
	ErrNewsTitleResource = errors.New("(NewsTitle) not found or empty")
	ErrCategoryNotFound  = errors.New("category not found")
	ErrDuplicateData     = errors.New("duplicate data")

	//Modular
	BaseResponseMessageSuccess = "success"
	BaseResponseMessageFailed  = "something not right"
)
