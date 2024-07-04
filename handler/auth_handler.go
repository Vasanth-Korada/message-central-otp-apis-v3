package handler

import (
	"message-central-v3/auth"
	"message-central-v3/service"

	"github.com/gofiber/fiber/v2"
)

func GenerateToken(c *fiber.Ctx) error {
	token, err := service.GenerateToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	auth.SetAuthToken(token)
	return c.JSON(fiber.Map{
		"token": token,
	})
}
