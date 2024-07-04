package handler

import (
	"message-central-v3/models"
	"message-central-v3/service"

	"github.com/gofiber/fiber/v2"
)

func SendOTP(c *fiber.Ctx) error {
	var request models.SendOTPRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	response, err := service.SendOTP(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(response)
}

func ValidateOTP(c *fiber.Ctx) error {
	var request models.ValidateOTPRequest
	if err := c.QueryParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid query parameters",
		})
	}

	response, err := service.ValidateOTP(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(response)
}
