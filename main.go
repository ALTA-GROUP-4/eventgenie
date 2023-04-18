package main

import (
	"github.com/labstack/echo/v4"
	userhandler "github.com/mujahxd/eventgenie/app/features/user/handler"
	userrepo "github.com/mujahxd/eventgenie/app/features/user/repository"
	userusecase "github.com/mujahxd/eventgenie/app/features/user/usecase"
	"github.com/mujahxd/eventgenie/config"
	"github.com/mujahxd/eventgenie/routes"
	"github.com/mujahxd/eventgenie/utils/database"
)

func main() {
	e := echo.New()
	loadConfig := config.InitConfig()
	db := database.ConnectionDB(loadConfig)

	// database
	database.Migrate(db)
	userRepository := userrepo.NewRepository(db)

	userService := userusecase.NewService(userRepository)

	userHandler := userhandler.NewHandler(userService)

	routes.UserRoute(e, userHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
