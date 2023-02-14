package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func GetUser(c *fiber.Ctx) error {
	return c.JSON(User{
		Id:   1,
		Name: "André",
		Age:  10,
	})
}
