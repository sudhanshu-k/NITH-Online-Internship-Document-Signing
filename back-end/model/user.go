package model

import (
	"time"

	guuid "github.com/google/uuid"
)

// store user details
type User struct {
	ID        guuid.UUID `json:"-"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname,omitempty"`
	Email     string     `json:"email"`
	IsFaculty bool       `json:"isfaculty"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"-" `
	UpdatedAt time.Time  `json:"-"`
}

// store user details to return to api caller
type UserResponse struct {
	ID        guuid.UUID `json:"-"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname,omitempty"`
	Email     string     `json:"email"`
	IsFaculty bool       `json:"isfaculty"`
	Level     string     `json:"level"`
	IsLog     bool       `json:"isloggedin"`
}

// convert user->userResponse
func FilterUserRecord(user *User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		IsFaculty: user.IsFaculty,
	}
}
