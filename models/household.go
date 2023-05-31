package models

import "gorm.io/gorm"

type Household struct {
	gorm.Model

	Name string `json:"name"`
	Residents []Resident
}