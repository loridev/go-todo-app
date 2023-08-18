package models

import "time"

type Todo struct {
	id        uint   `gorm:"primary_key"`
	Title     string `gorm:"not null"`
	Done      bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
