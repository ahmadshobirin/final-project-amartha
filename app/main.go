package main

import (
	"log"
	"time"

	_cityUsecase "main-backend/bussiness/city"
	_cityController "main-backend/controller/city"
	_cityRepo "main-backend/driver/database/city"

	_roleUsecase "main-backend/bussiness/role"
	_roleController "main-backend/controller/role"
	_roleRepo "main-backend/driver/database/role"

	_userUsecase "main-backend/bussiness/user"
	_userController "main-backend/controller/user"
	_userRepo "main-backend/driver/database/user"

	_clinicUsecase "main-backend/bussiness/clinic"
	_clinicController "main-backend/controller/clinic"
	_clinicRepo "main-backend/driver/database/clinic"

	_queueUsecase "main-backend/bussiness/queue"
	_queueController "main-backend/controller/queue"
	_queueRepo "main-backend/driver/database/queue"

	_authUsecase "main-backend/bussiness/auth"
	_authController "main-backend/controller/auth"

	_dbHelper "main-backend/driver/mysql"

	"main-backend/app/middleware"
	_routes "main-backend/app/routers"

	"github.com/labstack/echo/v4"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	configdb := _dbHelper.ConfigDB{
		DBUsername: viper.GetString(`database.user`),
		DBPassword: viper.GetString(`database.pass`),
		DBHost:     viper.GetString(`database.host`),
		DBPort:     viper.GetString(`database.port`),
		DBDatabase: viper.GetString(`database.name`),
	}
	db := configdb.InitialDB()

	configJWT := middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	cityRepo := _cityRepo.NewCityRepository(db)
	cityUsecase := _cityUsecase.NewCityUsecase(timeoutContext, cityRepo)
	cityCtrl := _cityController.NewCityController(e, cityUsecase)

	roleRepo := _roleRepo.NewRoleRepository(db)
	roleUsecase := _roleUsecase.NewRoleUsecase(timeoutContext, roleRepo)
	roleCtrl := _roleController.NewRoleController(e, roleUsecase)

	userRepo := _userRepo.NewUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(timeoutContext, userRepo)
	userCtrl := _userController.NewUserController(e, userUsecase, &configJWT)

	clinicRepo := _clinicRepo.NewClinicRepository(db)
	clinicUsecase := _clinicUsecase.NewClinicUsecase(timeoutContext, clinicRepo)
	clinicCtrl := _clinicController.NewClinicController(e, clinicUsecase)

	authUsecase := _authUsecase.NewAuthUsecase(timeoutContext, userUsecase, roleUsecase, &configJWT)
	authCtrl := _authController.NewAuthController(e, authUsecase)

	queueRepo := _queueRepo.NewQueueRepository(db)
	queueUsecase := _queueUsecase.NewQueueUsecase(timeoutContext, queueRepo)
	queueCtrl := _queueController.NewQueueController(e, queueUsecase, &configJWT)

	routesInit := _routes.ControllerList{
		JWTMiddleware:    configJWT.Init(),
		CityController:   *cityCtrl,
		RoleController:   *roleCtrl,
		UserController:   *userCtrl,
		AuthController:   *authCtrl,
		ClinicController: *clinicCtrl,
		QueueController:  *queueCtrl,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))

}
