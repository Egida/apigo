package controller

import (
	"api/model"
	"api/pdns"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joeig/go-powerdns/v3"
	"github.com/spf13/viper"
)

func AddZone(c *fiber.Ctx) error {

	domain := c.Params("domain")

	pwdns := powerdns.NewClient(viper.GetString("app.powerdnsserver"), "localhost", map[string]string{"X-API-Key": viper.GetString("app.powerdnskey")}, nil)
	ctx := context.Background()

	zone, err := pwdns.Zones.AddMaster(ctx, domain, true, "", false, "foo", "foo", true, []string{"ns1.dnic.icu.", "ns2.dnic.icu."})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
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
		"result":  zone,
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
