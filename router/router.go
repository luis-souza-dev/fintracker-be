package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luis-souza-dev/fintracker-be/handler"
)



func InitRoutes(app *fiber.App) {
	household := app.Group("/household")

	household.Get("/", handler.GetHousehold)
	household.Put("/", handler.EditHousehold)

	resident := app.Group("/residents")

	resident.Get("/", handler.GetResident)
	resident.Get("/:id", handler.GetResident)
	resident.Delete("/:id", handler.DeleteResident)
	resident.Post("/", handler.CreateResident)
	resident.Put("/", handler.EditResident)

	expensesCategories := app.Group("/expenses-categories")
	expensesCategories.Get("/", handler.GetExpensesCategories)
	expensesCategories.Post("/", handler.CreateExpensesCategories)
	expensesCategories.Delete("/:id", handler.DeleteExpensesCategories)

	expenses := app.Group("expenses")

	expenses.Get("/", handler.GetExpense)
	expenses.Get("/notes/:id", handler.GetExpenseNotes)
	expenses.Delete("/:id", handler.DeleteExpense)
	expenses.Post("/", handler.CreateExpense)
	expenses.Put("/", handler.EditExpense)

}