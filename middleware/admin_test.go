package middleware

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"api/model"
)

func TestRequireRole(t *testing.T) {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		role := c.Get("Role")
		c.Locals("user", model.User{Role: model.Role(role)})
		return c.Next()
	})
	app.Use(RequireRole(model.RolePro))
	app.Get("/protected", func(c *fiber.Ctx) error { return c.SendString("OK") })

	tests := []struct {
		name     string
		role     model.Role
		wantCode int
	}{
		{
			"admin have access",
			model.RoleAdmin,
			fiber.StatusOK,
		},
		{
			"pro have access",
			model.RolePro,
			fiber.StatusOK,
		},
		{
			"standard doesn't have access",
			model.RoleStandard,
			fiber.StatusForbidden,
		},
		{
			"unknown role doesn't have access",
			"foobar-role",
			fiber.StatusForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(fiber.MethodGet, "/protected", nil)
			req.Header.Set("Role", string(tt.role))

			resp, _ := app.Test(req, -1)

			assert.Equal(t, tt.wantCode, resp.StatusCode)
		})
	}
}
