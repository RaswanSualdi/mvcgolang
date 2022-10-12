package model

import (
	"time"
)

type User struct {
	// gorm.Model
	// CreditCards    []Letter
	ID             int
	Name           string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
