package repository

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name         string
	Date         datatypes.Date
	HostedBy     string
	Place        string
	TypeOfTicket bool
	Price        int
	GoalQuota    int
	CurrentQuota int
	ImageURL     string
	Description  string
	UserID       uint
}
