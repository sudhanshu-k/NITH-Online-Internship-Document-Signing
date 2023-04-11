package model

import (
	"time"

	guuid "github.com/google/uuid"
)

type Student struct {
	ID        guuid.UUID `gorm:"type:uuid" json:"-"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname,omitempty"`
	Email     string     `gorm:"unique" json:"email"`
	Password  string     `json:"password"`
	Sessions  []Session  `gorm:"foreignKey:UserRefer; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
	// Products  []Product  `gorm:"foreignKey:UserRefer; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-" `
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
}

func FilterUserRecord(user *Student) Student {
	return Student{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
