package model

import (
	"time"

	guuid "github.com/google/uuid"
)

type UGInternForm struct {
	ID            guuid.UUID `json:"-"`
	Name          string     `json:"name"`
	FatherName    string     `json:"fathername"`
	Address       string     `json:"address"`
	Contact       string     `json:"contact"`
	CompanyName   string     `json:"companyname"`
	Email         string     `json:"email"`
	AreaOfIntrest string     `json:"aoi"`
	IsOffline     bool       `json:"isoffline"`
	StartDay      time.Time  `json:"startday"`
	EndDay        time.Time  `json:"endday"`
	Weeks         int        `json:"weeks"`
	FromTPO       bool       `josn:"formtpo"`
	Stipend       int        `josn:"stipend"`
	FormDate      time.Time  `json:"fromdate"`
	RemarksDept   string     `json:"remarksdept"`
	RemarksFI     string     `json:"remarksfi"`
	Level         int        `json:"level"`
	CreatedAt     time.Time  `json:"-" `
	UpdatedAt     time.Time  `json:"-"`
}

type user struct {
	Name       string `json:"name"`
	RollNumber string `json:"roll"`
}

type FormResponse struct {
	ID       guuid.UUID `json:"UUID"`
	FormType string     `json:"FormName"`
	User     user       `json:"user"`
}
