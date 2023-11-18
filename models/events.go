package models

import "github.com/jinzhu/gorm"

type Event struct {
	gorm.Model
	UserID      uint
	Title       string
	Description string
	Duration    int // Duration in minutes
	// Other relevant fields for the event
}
