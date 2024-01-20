package middleware

import (
	"api/helper"

	"github.com/gofiber/fiber/v2"
)

func BindCurrentUser(c *fiber.Ctx) error {
	user, err := helper.CurrentUser(c)

	if err == nil {
		c.Locals("user", user)
		c.Bind(fiber.Map{"User": user})
	}

	return c.Next()
}
