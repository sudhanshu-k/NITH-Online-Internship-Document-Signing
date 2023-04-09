package model

import (
	"time"

	"github.com/google/uuid"
)

type Faculty struct {
	ID        uuid.UUID `gorm:"type:uuid" json:"-"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname,omitempty"`
	Email     string    `json:"email"`
	Department	string	`json:"department,omitempty"`
	Password  string    `json:"password"`
	Level     int8      `json:"level"`
	Sessions  []Session  `gorm:"foreignKey:UserRefer; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
	// Products  []Product  `gorm:"foreignKey:UserRefer; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-" `
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"-"`
}
