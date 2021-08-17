package routers

import (
	"main-backend/controller/auth"
	"main-backend/controller/city"
	"main-backend/controller/role"
	"main-backend/controller/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	CityController city.CityController
	RoleController role.RoleController
	UserController user.UserController
	AuthController auth.AuthController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	r := e.Group("/api/v1")

	authRouter := r.Group("/auth")
	authRouter.POST("/login", cl.AuthController.Login)
	authRouter.POST("/register", cl.AuthController.Register)

	r.Use(middleware.JWTWithConfig(cl.JWTMiddleware))

	cityRouter := r.Group("/city")
	cityRouter.GET("", cl.CityController.Find)
	cityRouter.GET("/:id", cl.CityController.FindByID)

	roleRouter := r.Group("/role")
	roleRouter.GET("", cl.RoleController.Find)
	roleRouter.GET("/:id", cl.RoleController.FindByID)

	userRouter := r.Group("/user")
	userRouter.POST("", cl.UserController.Store)
}
