package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mujahxd/eventgenie/app/features/user"
)

func UserRoute(e *echo.Echo, h user.UserHandler) {
	// e.Pre(middleware.RemoveTrailingSlash())
	// e.Use(middleware.CORS())
	// e.Use(middleware.Logger())

	e.POST("/users", h.RegisterUser())

}
