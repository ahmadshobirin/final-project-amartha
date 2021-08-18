package routers

import (
	"main-backend/controller/auth"
	"main-backend/controller/city"
	"main-backend/controller/clinic"
	"main-backend/controller/queue"
	"main-backend/controller/role"
	"main-backend/controller/user"

	"main-backend/app/middleware"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	JWTMiddleware    *middleware.ConfigJWT
	CityController   city.CityController
	RoleController   role.RoleController
	UserController   user.UserController
	AuthController   auth.AuthController
	ClinicController clinic.ClinicController
	QueueController  queue.QueueController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	amMiddleware := *cl.JWTMiddleware
	amMiddleware.Role = []string{"AM"}
	usMiddleware := *cl.JWTMiddleware
	usMiddleware.Role = []string{"AM"}

	usersMiddleware := *cl.JWTMiddleware
	usersMiddleware.Role = []string{"AM", "US", "SA"}

	r := e.Group("/api/v1")

	authRouter := r.Group("/auth")
	authRouter.POST("/login", cl.AuthController.Login)
	authRouter.POST("/register", cl.AuthController.Register)

	cityRouter := r.Group("/city", amMiddleware.VerifyRole)
	cityRouter.GET("", cl.CityController.GetAll)
	cityRouter.GET("/:id", cl.CityController.FindByID)
	cityRouter.POST("", cl.CityController.Store)
	cityRouter.PUT("/:id", cl.CityController.Update)
	cityRouter.DELETE("/:id", cl.CityController.Delete)

	roleRouter := r.Group("/role", amMiddleware.VerifyRole)
	roleRouter.GET("", cl.RoleController.Find)
	roleRouter.GET("/:id", cl.RoleController.FindByID)

	userRouter := r.Group("/user", usersMiddleware.VerifyRole)
	userRouter.GET("", cl.UserController.Fetch)
	userRouter.GET("/profile", cl.UserController.Profile)
	userRouter.PUT("/profile", cl.UserController.Update)
	userRouter.POST("", cl.UserController.Store)

	clinicRouter := r.Group("/clinic", usersMiddleware.VerifyRole)
	clinicRouter.GET("", cl.ClinicController.Fetch)
	clinicRouter.GET("/:id", cl.ClinicController.FindByID)
	clinicRouter.POST("", cl.ClinicController.Store)
	clinicRouter.PUT("/:id", cl.ClinicController.Update)
	clinicRouter.DELETE("/:id", cl.ClinicController.Delete)

	transactionRouter := r.Group("/transaction")
	transactionRouter.POST("", cl.QueueController.Store)
	transactionRouter.PUT("/:id", cl.QueueController.Update)

}
