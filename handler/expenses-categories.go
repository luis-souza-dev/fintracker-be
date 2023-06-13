package handler

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/luis-souza-dev/fintracker-be/database"
	"github.com/luis-souza-dev/fintracker-be/models"
)

type ExpensesCategoriesQuery struct {
	Name string `query:"name"`
}

func GetExpensesCategories(c *fiber.Ctx) error {
	result := &[]models.ExpensesCategories{}
	query := &ExpensesCategoriesQuery{}
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(fmt.Sprintf("query err%v", err))
	}
	
	database.DB.Where(query).Find(result)

	jsonval, err := json.Marshal(result);
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(fmt.Sprintf("json err %v", err))
	}
	return c.Status(200).SendString(string(jsonval))

}

func CreateExpensesCategories(c *fiber.Ctx) error {
	expenseCategory := &models.ExpensesCategories{}
	if err := c.BodyParser(expenseCategory); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(fmt.Sprintf("body err%v", err))
	}
	if expenseCategory.Name == "" {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Missing name")
	}
	res := database.DB.Create(expenseCategory)
	if res.Error != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(fmt.Sprintf("creation err%v", res.Error))
	}

	return c.Status(fiber.StatusCreated).SendString(fmt.Sprint(expenseCategory.ID))
}

func DeleteExpensesCategories(c *fiber.Ctx) error {
	id := c.Params("id");
	if id == "" {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Missing id")
	}
	res := database.DB.Delete(&models.ExpensesCategories{}, id)

	if res.Error != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Deletion err")
	}
	return c.Status(200).SendString("done")
}