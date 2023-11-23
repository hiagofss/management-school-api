package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name     string `json:"name" validate:"nonzero, regexp=^[a-zA-Z ]*$"`
	Email    string `json:"email" gorm:"unique" validate:"regexp=^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$"`
	Document string `json:"document" validate:"len=11, regexp=^[0-9]*$"`
	// Birthdate *time.Time `json:"birthdate"`
	// Grade     int `json:"grade"`
}

func ValidateStudent(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}

	return nil
}
