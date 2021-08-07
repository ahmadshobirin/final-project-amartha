package main

import (
	"main-backend/config"
	"main-backend/routers"

	"github.com/joho/godotenv"
)

func main() {
	// load .env
	godotenv.Load()

	// initDB
	initDB := config.InitDB()

	// migrate table
	config.MigrateTables(initDB)

	routers.Api()
}
