package controller

import (
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/gofiber/fiber/v2"

	"api/helper"
	"api/model"
)

func IndexView(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func LoginView(c *fiber.Ctx) error {
	if c.Locals("user") != nil {
		return c.Redirect("/admin")
	}

	return c.Render("login", fiber.Map{})
}

func LoginForm(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := model.FindUserByUsername(username)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := user.ValidatePassword(password); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    jwt,
		HTTPOnly: true,
		Expires:  time.Now().Add(24 * 7 * time.Hour),
	})

	return c.Redirect("/admin")
}

func LogoutForm(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		HTTPOnly: true,
		Expires:  time.Now(),
	})

	return c.Redirect("/admin")
}

func AccountsIndex(c *fiber.Ctx) error {
	if c.Locals("user") == nil {
		return c.Redirect("/admin/login")
	}

	users, err := model.ListUsers()
	if err != nil {
		return err
	}

	return c.Render("accounts", fiber.Map{
		"Users": users,
		"Roles": []string{"standard", "pro", "admin"},
	})
}

func AccountCreateForm(c *fiber.Ctx) error {
	ips := c.FormValue("ips")
	user := model.User{
		Username: c.FormValue("username"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
		Role:     model.Role(c.FormValue("role")),
	}

	_, err := user.Save()
	if err != nil {
		return err
	}

	if ips != "" {
		var s []string
		if err := jsoniter.Unmarshal([]byte(ips), &s); err != nil {
			return err
		}

		for _, ip := range s {
			model.CreateIP(user, ip)
		}
	}

	return c.Redirect("/admin/accounts")
}

func AccountEditForm(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	user, err := model.FindUserById(uint(id))
	if err != nil {
		return err
	}

	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	role := c.FormValue("role")
	ips := c.FormValue("ips")

	if err := user.Update(username, email, password, model.Role(role)); err != nil {
		return err
	}

	if ips != "" {
		var s []string
		if err := jsoniter.Unmarshal([]byte(ips), &s); err != nil {
			return err
		}

		for _, ip := range s {
			model.CreateIP(user, ip)
		}
	}

	return c.Redirect("/admin/accounts")
}

func AccountDeleteForm(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	user, err := model.FindUserById(uint(id))
	if err != nil {
		return err
	}

	if err := user.Delete(); err != nil {
		return err
	}

	return c.Redirect("/admin/accounts")
}

func LogsView(c *fiber.Ctx) error {
	return c.Render("logs", fiber.Map{})
}

func AccountRevoke(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	user, err := model.FindUserById(uint(id))
	if err != nil {
		return err
	}

	if err := user.RevokeTokens(); err != nil {
		return err
	}

	return c.Redirect("/admin/accounts")
}

func CreateToken(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	user, err := model.FindUserById(uint(id))
	if err != nil {
		return err
	}

	apiKey, err := model.CreateAPIKey(user)
	if err != nil {
		return err
	}

	return c.Render("token", fiber.Map{
		"User":  user,
		"Token": apiKey.Token,
	})
}
