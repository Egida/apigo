package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"api/model"
)

func RequireRole(role model.Role) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user")
		if user == nil {
			return fiber.NewError(fiber.StatusForbidden, fmt.Sprintf("you must have %s role to perform this action", role))
		}

		userRole := user.(model.User).Role

		val, ok := model.Roles[userRole]
		if !ok {
			return fiber.NewError(fiber.StatusForbidden, fmt.Sprintf("you must have %s role to perform this action", role))
		}

		if val > model.Roles[role] {
			return fiber.NewError(fiber.StatusForbidden, fmt.Sprintf("you must have %s role to perform this action", role))
		}

		return c.Next()
	}
}

func RequireAdmin(c *fiber.Ctx) error {
	return RequireRole(model.RoleAdmin)(c)
}

func RequirePro(c *fiber.Ctx) error {
	return RequireRole(model.RolePro)(c)
}
