package handler

import "github.com/gofiber/fiber/v2"

func GetHousehold(c *fiber.Ctx) error {
	return c.Status(200).SendString("hey");
}

func EditHousehold(c *fiber.Ctx) error {
	return c.Status(200).SendString("hey2");
}