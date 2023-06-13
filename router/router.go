package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luis-souza-dev/fintracker-be/handler"
)



func InitRoutes(app *fiber.App) {

	resident := app.Group("/residents")

	resident.Get("/", handler.GetResident)
	resident.Get("/:id", handler.GetResident)
	resident.Delete("/:id", handler.DeleteResident)
	resident.Post("/", handler.CreateResident)
	resident.Put("/:id", handler.EditResident)

	expensesCategories := app.Group("/expenses-categories")
	expensesCategories.Get("/", handler.GetExpensesCategories)
	expensesCategories.Post("/", handler.CreateExpensesCategories)
	expensesCategories.Delete("/:id", handler.DeleteExpensesCategories)

	expenses := app.Group("expenses")

	expenses.Get("/", handler.GetExpense)
	expenses.Delete("/:id", handler.DeleteExpense)
	expenses.Post("/", handler.CreateExpense)
	expenses.Put("/:id", handler.EditExpense)

}