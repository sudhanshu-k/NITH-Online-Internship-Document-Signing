package model

import (
	"time"

	"github.com/google/uuid"
)

type Faculty struct {
	ID        uuid.UUID `gorm:"type:uuid" json:"-"`
	FirstName string    `json:"firstname"`
	LastName string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Level     int8      `json:"-"`
	// Sessions  []Session  `gorm:"foreignKey:UserRefer; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
	// Products  []Product  `gorm:"foreignKey:UserRefer; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-" `
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"-"`
}
