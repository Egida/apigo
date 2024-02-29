package controller

import (
	"api/model"
	"api/pdns"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CZone(c *fiber.Ctx) error {
	var input model.AddZoneInput

	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := model.Validate.Struct(&input); err != nil {
		return err
	}
	if input.Kind == "" {
		input.Kind = "Master"
	}

	if input.NameServers == nil || len(input.NameServers) == 0 {
		input.NameServers = []string{"ns1.dnic.icu.", "ns2.dnic.icu."}
	}

	zone, err := pdns.Add(input)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "Zone wurde erstellt",
		"result":  zone,
	})
}
