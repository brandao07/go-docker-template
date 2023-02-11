package main

import (
	"github.com/brandao07/go-docker-template/cmd/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
}
