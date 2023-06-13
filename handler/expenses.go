package handler

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/luis-souza-dev/fintracker-be/database"
	"github.com/luis-souza-dev/fintracker-be/models"
)

func CreateExpense(c *fiber.Ctx) error {

	expense := new(models.Expense)
	
	if err := c.BodyParser(expense); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(fmt.Sprintf("body err%v", err))
	}

	if expense.Name == "" {
		return c.Status(fiber.ErrBadRequest.Code).SendString("missing name")
	}

	noCat := models.ExpensesCategories{}
	if expense.ExpensesCategoriesID == 0 && expense.ExpensesCategories == noCat {
		return c.Status(fiber.ErrBadRequest.Code).SendString("missing category")
	}

	if expense.Date == "" {
		return c.Status(fiber.ErrBadRequest.Code).SendString("missing date")
	}

	if expense.Total == 0 {
		return c.Status(fiber.ErrBadRequest.Code).SendString("can not enter negative values")
	}

	if expense.Status == "" {
		expense.Status = "Paid"
	}

	res := database.DB.Create(expense)
	if res.Error != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("smthg went wrong")
	}

	return c.Status(200).SendString(fmt.Sprint(expense.ID))
}

func GetExpense(c *fiber.Ctx) error {
	result := &[]models.Expense{}
	query := &models.Expense{}
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
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Missing id")
	}

	expense := new(models.Expense)
	reqBody := new(models.Expense)

	if err := c.BodyParser(reqBody); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(fmt.Sprintf("body err%v", err))
	}

	intId, err := strconv.Atoi(id)

	if err == nil {
		expense.ID = uint(intId)
		res := database.DB.Model(expense).Updates(reqBody)

		if res.Error == nil {
			return c.Status(200).SendString(fmt.Sprintf("body err %v, \n and queryBody %v", expense, reqBody))
		} else {
			return c.Status(fiber.ErrBadRequest.Code).SendString("Error while updating")
		}
	} else {
		return c.Status(fiber.ErrBadRequest.Code).SendString(fmt.Sprintf("incorrect id value: %v", err))
	}
}

func DeleteExpense(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Missing id")
	}
	res := database.DB.Delete(&models.Expense{}, id)

	if res.Error != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Deletion err")
	}
	return c.Status(200).SendString("done")
}

