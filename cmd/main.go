package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	setupRoutes(app)
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
