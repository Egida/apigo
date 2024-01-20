package middleware

import (
	"api/helper"

	"github.com/gofiber/fiber/v2"
)

func JWTAuthMiddleware(c *fiber.Ctx) error {
	err := helper.ValidateJWT(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	return c.Next()
}
