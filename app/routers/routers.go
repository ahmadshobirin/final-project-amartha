package routers

import (
	"main-backend/controller/auth"
	"main-backend/controller/city"
	"main-backend/controller/clinic"
	"main-backend/controller/queue"
	"main-backend/controller/role"
	"main-backend/controller/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware    middleware.JWTConfig
	CityController   city.CityController
	RoleController   role.RoleController
	UserController   user.UserController
	AuthController   auth.AuthController
	ClinicController clinic.ClinicController
	QueueController  queue.QueueController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	r := e.Group("/api/v1")

	authRouter := r.Group("/auth")
	authRouter.POST("/login", cl.AuthController.Login)
	authRouter.POST("/register", cl.AuthController.Register)

	r.Use(middleware.JWTWithConfig(cl.JWTMiddleware))

	cityRouter := r.Group("/city")
	cityRouter.GET("", cl.CityController.GetAll)
	cityRouter.GET("/:id", cl.CityController.FindByID)
	cityRouter.POST("", cl.CityController.Store)
	cityRouter.PUT("/:id", cl.CityController.Update)
	cityRouter.DELETE("/:id", cl.CityController.Delete)

	roleRouter := r.Group("/role")
	roleRouter.GET("", cl.RoleController.Find)
	roleRouter.GET("/:id", cl.RoleController.FindByID)

	userRouter := r.Group("/user")
	userRouter.GET("", cl.UserController.Fetch)
	userRouter.GET("/profile", cl.UserController.Profile)
	userRouter.PUT("/profile", cl.UserController.Update)
	userRouter.POST("", cl.UserController.Store)

	clinicRouter := r.Group("/clinic")
	clinicRouter.GET("", cl.ClinicController.Fetch)
	clinicRouter.GET("/:id", cl.ClinicController.FindByID)
	clinicRouter.POST("", cl.ClinicController.Store)
	clinicRouter.PUT("/:id", cl.ClinicController.Update)
	clinicRouter.DELETE("/:id", cl.ClinicController.Delete)

	transactionRouter := r.Group("/transaction")
	transactionRouter.POST("", cl.QueueController.Store)
	transactionRouter.PUT("/:id", cl.QueueController.Update)

}
