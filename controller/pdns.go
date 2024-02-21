package controller

import (
	"api/model"
	"api/pdns"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AddZone(c *fiber.Ctx) error {
	var input model.AddZoneInput
	domain := c.Params("domain")
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

	zone, err := pdns.AddZone(input)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "Zone wurde erstellt",
		"result":  zone,
	})
}

func AddRecord(c *fiber.Ctx) error {
	domain := c.Params("domain")

	var input model.AddRecodInput
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := model.Validate.Struct(&input); err != nil {
		return err
	}

	err := pdns.AddRRSets(domain, input)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	zone, err := pdns.GetZone(domain)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "Record wurde erstellt",
		"result":  zone.RRsets,
	})
}
func RemoveZone(c *fiber.Ctx) error {
	domain := c.Params("domain")

	err := pdns.RemoveZone(domain)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Zone wurde gelöscht",
	})
}
