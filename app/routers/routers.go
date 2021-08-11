package routers

import (
	"main-backend/controller/city"

	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	CityController city.CityController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	r := e.Group("/api/v1")

	cityRouter := r.Group("/city")
	cityRouter.GET("", cl.CityController.Find)
	cityRouter.GET("/:id", cl.CityController.FindByID)
}
