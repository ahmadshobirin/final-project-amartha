package main

import (
	"log"
	"time"

	_cityUsecase "main-backend/bussiness/city"
	_cityController "main-backend/controller/city"
	_cityRepo "main-backend/driver/database/city"

	_dbHelper "main-backend/helper/database"

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

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	cityRepo := _cityRepo.NewCityRepository(db)
	cityUsecase := _cityUsecase.NewCityUsecase(timeoutContext, cityRepo)
	cityCtrl := _cityController.NewCityController(e, cityUsecase)

	routesInit := _routes.ControllerList{
		CityController: *cityCtrl,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))

}
