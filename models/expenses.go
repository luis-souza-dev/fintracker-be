package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model

	Name string `json:"name"`
	Total float32 `json:"total"`
	Recurrent bool `json:"recurrent"`
	Date string `json:"date"`
	Status string `json:"status"`
	ExpensesCategoriesID int
	ExpensesCategories ExpensesCategories
	ResidentID int
	Resident Resident
}