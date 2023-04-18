package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mujahxd/eventgenie/app/features/user"
	"github.com/mujahxd/eventgenie/helper"
)

type userHandler struct {
	userService user.Service
}

func NewHandler(userService user.Service) user.UserHandler {
	return &userHandler{userService}
}

func (uh *userHandler) RegisterUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input user.RegisterUserInput
		err := c.Bind(&input)
		if err != nil {
			c.Logger().Error(err.Error())
			response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", err.Error())
			return c.JSON(http.StatusUnprocessableEntity, response)
		}
		// validas inputan kosong
		validate := validator.New()
		err = validate.Struct(input)
		if err != nil {
			c.Logger().Error(err.Error())
			response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", err.Error())
			return c.JSON(http.StatusBadRequest, response)
		}

		newUser, err := uh.userService.RegisterUser(input)
		if err != nil {
			c.Logger().Error(err.Error())
			response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", err.Error())
			return c.JSON(http.StatusBadRequest, response)

		}

		response := helper.APIResponse("Account has been registered", http.StatusCreated, "success", newUser)
		return c.JSON(http.StatusCreated, response)
	}
}
