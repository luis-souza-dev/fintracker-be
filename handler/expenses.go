package handler

import "github.com/gofiber/fiber/v2"

func CreateExpense(c *fiber.Ctx) error {
	return c.Status(200).SendString("expenses works")
}

func GetExpense(c *fiber.Ctx) error {
	return c.Status(200).SendString("expenses works")
}

func GetExpenseNotes(c *fiber.Ctx) error {
	return c.Status(200).SendString("expenses works")
}

func EditExpense(c *fiber.Ctx) error {
	return c.Status(200).SendString("expenses works")
}

func DeleteExpense(c *fiber.Ctx) error {
	return c.Status(200).SendString("expenses works")
}

