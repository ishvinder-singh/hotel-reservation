package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ishvinder-singh/hotel-reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Alex",
		LastName:  "Smith",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("user")
}
