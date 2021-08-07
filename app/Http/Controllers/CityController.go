package controllers

import (
	"main-backend/app/models"
	"main-backend/config"
	"main-backend/helpers"

	"net/http"

	"github.com/labstack/echo/v4"
)

func CityIndex(c echo.Context) error {
	model := []models.City{}
	res := helpers.BaseResponse{}

	db := config.InitDB().Find(&model)
	if db.Error != nil {
		res = helpers.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: helpers.MessageFailed,
			Data:    nil,
		}

		return c.JSON(http.StatusOK, res)
	}

	res = helpers.BaseResponse{
		Code:    http.StatusOK,
		Message: helpers.MessageSuccess,
		Data:    model,
	}

	return c.JSON(http.StatusOK, res)

}
