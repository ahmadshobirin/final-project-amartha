package controller

import (
	"main-backend/helper/messages"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Total    int      `json:"total,omiempty"`
		Status   int      `json:"status"`
		Message  string   `json:"message"`
		Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccessResponseWithTotal(c echo.Context, param interface{}, total int) error {
	response := BaseResponse{}
	response.Meta.Total = total
	response.Meta.Status = http.StatusOK
	response.Meta.Message = messages.BaseResponseMessageSuccess
	response.Data = param

	return c.JSON(http.StatusOK, response)
}

func NewSuccessResponse(c echo.Context, param interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = messages.BaseResponseMessageSuccess
	response.Data = param

	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = messages.BaseResponseMessageFailed
	response.Meta.Messages = []string{err.Error()}

	return c.JSON(http.StatusOK, response)
}
