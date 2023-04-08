package model

import (
	"time"

	guuid "github.com/google/uuid"
)

type Profile struct {
	ID        guuid.UUID `gorm:"primaryKey" json:"-"`
	FirstName string     `json:"firstname"`
	LasttName string     `json:"lastname"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Level     int8       `json:"level"`
	// Sessions  []Session  `gorm:"foreignKey:UserRefer; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
	// Products  []Product  `gorm:"foreignKey:UserRefer; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"-" `
	UpdatedAt time.Time  `gorm:"autoUpdateTime:milli" json:"-"`
}
