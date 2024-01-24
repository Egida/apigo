package controller

import (
	"api/model"

	"github.com/gofiber/fiber/v2"
)

func ShowUser(c *fiber.Ctx) error {
	head := c.GetReqHeaders()
	token := head["X-Apikey"]
	isuser, err := model.FindAPIKey(token)
	if err != nil {
		return fiber.NewError(fiber.StatusOK, err.Error())
	}

	user, err := model.FindUserByUsername(isuser.User.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": map[string]string{
			"username": user.Username,
			"email":    user.Email,
			"role":     string(user.Role),
		},
	})
}
func ChangePassword(c *fiber.Ctx) error {
	head := c.GetReqHeaders()
	token := head["X-Apikey"]
	isuser, err := model.FindAPIKey(token)
	if err != nil {
		return fiber.NewError(fiber.StatusOK, err.Error())
	}
	password := c.FormValue("password")
	user, err := model.FindUserByUsername(isuser.User.Username)
	if err := user.Update(user.Username, user.Email, password, user.Role); err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    "Password changed successfully",
	})
}
