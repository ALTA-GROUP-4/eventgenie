package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mujahxd/eventgenie/config"
	"github.com/mujahxd/eventgenie/utils/database"
)

func main() {
	e := echo.New()
	loadConfig := config.InitConfig()
	db := database.ConnectionDB(loadConfig)

	// database
	database.Migrate(db)

	e.Logger.Fatal(e.Start(":8000"))
}
