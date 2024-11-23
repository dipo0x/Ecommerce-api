package controller

import (
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	var user = c.Locals("user")
	
	if user == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"error": "User not available"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":  user,
	})
}
