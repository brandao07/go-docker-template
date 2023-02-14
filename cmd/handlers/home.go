package handlers

import "github.com/gofiber/fiber/v2"

func GetHome(c *fiber.Ctx) error {
	return c.SendString("Hello Gophers!")
}

func DeleteHome(c *fiber.Ctx) error {
	return c.SendString("Delete Gophers!")
}
