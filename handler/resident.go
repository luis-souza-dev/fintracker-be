package handler

import "github.com/gofiber/fiber/v2"

func GetResident(c *fiber.Ctx) error {
	return c.Status(200).SendString("Resident works")
}

func CreateResident(c *fiber.Ctx) error {
	return c.Status(200).SendString("Resident works")
}

func EditResident(c *fiber.Ctx) error {
	return c.Status(200).SendString("Resident works")
}

func DeleteResident(c *fiber.Ctx) error {
	return c.Status(200).SendString("Resident works")
}
