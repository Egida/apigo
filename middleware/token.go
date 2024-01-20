package middleware

import (
	"api/model"

	"github.com/gofiber/fiber/v2"
)

func APIKeyAuthMiddleware(c *fiber.Ctx) error {
	head := c.GetReqHeaders()
	token := head["X-Apikey"]

	apiKey, err := model.FindAPIKey(token)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid or missing API key")
	}
	if apiKey == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid or missing API key")
	}

	c.Locals("user", apiKey.User)
	return c.Next()
}
