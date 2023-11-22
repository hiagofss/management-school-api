package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Document string `json:"document"`
	// Birthdate *time.Time `json:"birthdate"`
	// Grade     int `json:"grade"`
}
