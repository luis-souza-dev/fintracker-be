package models

import "gorm.io/gorm"

type Resident struct {
	gorm.Model

	Name string `json:"name"`
	Birthday string `json:"birthday"`
	Income float32 `json:"income"`
	Expenses []Expense
}