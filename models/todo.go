package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
  gorm.Model
	id        uint
	Title     string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Todo) TableName() string {
  return "todos"
}
