package handler

import "github.com/gofiber/fiber/v2"

func GetExpensesCategories(c *fiber.Ctx) error {
	return c.Status(200).SendString("expenses categories works")
}

func CreateExpensesCategories(c *fiber.Ctx) error {
	return c.Status(200).SendString("expenses categories works")
}

func DeleteExpensesCategories(c *fiber.Ctx) error {
	return c.Status(200).SendString("expenses categories works")
}