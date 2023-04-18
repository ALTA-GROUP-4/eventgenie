package repository

import (
	"github.com/mujahxd/eventgenie/app/features/user"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) user.Repository {
	return &repository{db}
}

func (r *repository) Save(user user.User) (user.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
