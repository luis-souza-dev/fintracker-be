package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/luis-souza-dev/fintracker-be/router"
)

func main() {
	app := fiber.New()

	router.InitRoutes(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(map[string]string{
			"foo": "bar",
		})
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})


	log.Fatal(app.Listen(":3000"))
}