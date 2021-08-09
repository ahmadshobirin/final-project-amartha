package routers

import (
	"os"

	controller "main-backend/app/Http/Controllers"

	"github.com/labstack/echo/v4"
)

func Api() {
	e := echo.New()

	e.GET("/city", controller.CityIndex)

	e.Logger.Fatal(e.Start(os.Getenv("APP_HOST")))
}
