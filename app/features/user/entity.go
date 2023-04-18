package user

import (
	"github.com/labstack/echo/v4"
	"github.com/mujahxd/eventgenie/app/features/event/repository"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string
	Email          string
	Password       string
	AvatarFileName string
	Events         []repository.Event
}

type Repository interface {
	Save(user User) (User, error)
}
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type UserHandler interface {
	RegisterUser() echo.HandlerFunc
}
