package model

import (
	"time"
)

// User has many CreditCards, UserID is the foreign key

type Letter struct {
	// gorm.Model
	ID            int
	Uuid          string
	MdLettersId   int
	MdCompaniesId int
	Letter        string
	Description   string
	Link          string
	Slug          string
	DateLetter    string
	YearLetter    string
	MonthLetter   string
	NoLetter      string
	Images        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
