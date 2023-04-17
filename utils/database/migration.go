package database

import (
	eventrepo "github.com/mujahxd/eventgenie/app/features/event/repository"
	userrepo "github.com/mujahxd/eventgenie/app/features/user/repository"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	users := userrepo.User{}
	events := eventrepo.Event{}

	db.AutoMigrate(users)
	db.AutoMigrate(events)
}
