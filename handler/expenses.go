package handler

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/luis-souza-dev/fintracker-be/database"
	"github.com/luis-souza-dev/fintracker-be/models"
)

func CreateExpense(c *fiber.Ctx) error {

	var expense models.Expenses
	if err := c.BodyParser(&expense); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(fmt.Sprintf("body err%v", err))
	}

	if expense.Name == "" {
		return c.Status(fiber.ErrBadRequest.Code).SendString("missing name")
	}

	if expense.Date == "" {
		return c.Status(fiber.ErrBadRequest.Code).SendString("missing date")
	}

	if expense.Total == 0 {
		return c.Status(fiber.ErrBadRequest.Code).SendString("can not enter negative values")
	}
	res := database.DB.Create(&expense)
	if res.Error != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("smthg went wrong")
	}
	var i int64
	res.Count(&i);

	return c.Status(200).SendString(fmt.Sprintf("expenses works! total:%d", i))
}

func GetExpense(c *fiber.Ctx) error {
	result := &[]models.Expenses{}
	query := &models.Expenses{}
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

func EditExpense(c *fiber.Ctx) error {
	return c.Status(200).SendString("expenses works")
}

func DeleteExpense(c *fiber.Ctx) error {
	id := c.Params("id");
	if id == "" {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Missing id")
	}
	res := database.DB.Delete(&models.Expenses{}, id)

	if res.Error != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Deletion err")
	}
	return c.Status(200).SendString("done")
}

