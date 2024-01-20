package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"api/model"
)

func RequireIP(c *fiber.Ctx) error {
	user := c.Locals("user")
	if user == nil {
		return fiber.NewError(fiber.StatusForbidden, "this ip is not whitelisted: user not found")
	}

	userID := user.(model.User).ID
	rip := c.IP()
	fmt.Println("user_id", userID, "ip", rip)

	allowed := model.IPAllowed(userID, rip)
	if !allowed {
		return fiber.NewError(fiber.StatusForbidden, fmt.Sprintf("this ip is not whitelisted: %s, %d", rip, userID))
	}

	return c.Next()
}
