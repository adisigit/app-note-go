package models

import (
	"github.com/google/uuid"
)

type Note struct {
	BaseModel
	Title   string    `gorm:"not null"`
	Content string    `gorm:"not null"`
	UserID  uuid.UUID `gorm:"not null"`
}
