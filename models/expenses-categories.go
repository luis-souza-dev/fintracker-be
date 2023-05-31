package models

import "gorm.io/gorm"

type ExpensesCategories struct {
	gorm.Model

	Name string `json:"name"`
}