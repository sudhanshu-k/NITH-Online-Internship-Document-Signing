package model

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Sessionid uuid.UUID `gorm:"primaryKey" json:"sessionid"`
	Expires   time.Time `json:"-"`
	UserRefer uuid.UUID `json:"-"`
	CreatedAt int64     `gorm:"autoCreateTime" json:"-" `
}
