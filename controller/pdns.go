package controller

import (
	"api/model"
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

	zone, err := pwdns.Zones.AddMaster(ctx, domain, true, "", false, "INCEPTION-INCREMENT", "DEFAULT", true, []string{"ns1.dnic.icu.", "ns2.dnic.icu."})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "Zone wurde erstellt",
		"result":  zone,
	})
}
func RemoveZone(c *fiber.Ctx) error {
	domain := c.Params("domain")
	pwdns := powerdns.NewClient(viper.GetString("app.powerdnsserver"), "localhost", map[string]string{"X-API-Key": viper.GetString("app.powerdnskey")}, nil)
	ctx := context.Background()
	err := pwdns.Zones.Delete(ctx, domain)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Zone wurde gelöscht",
	})
}

func AddRecord(c *fiber.Ctx) error {
	domain := c.Params("domain")
	pwdns := powerdns.NewClient(viper.GetString("app.powerdnsserver"), "localhost", map[string]string{"X-API-Key": viper.GetString("app.powerdnskey")}, nil)
	ctx := context.Background()
	var input model.RecordIn
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := model.Validate.Struct(&input); err != nil {
		return err
	}

	err := pwdns.Records.Add(ctx, domain, input.Name+"."+domain, powerdns.RRType(input.Type), uint32(input.TTL), []string{input.Data})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "Record wurde erstellt",
	})
}
func RemoveRecord(c *fiber.Ctx) error {
	domain := c.Params("domain")
	pwdns := powerdns.NewClient(viper.GetString("app.powerdnsserver"), "localhost", map[string]string{"X-API-Key": viper.GetString("app.powerdnskey")}, nil)
	ctx := context.Background()
	var input model.RecordDeletIn
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := model.Validate.Struct(&input); err != nil {
		return err
	}
	err := pwdns.Records.Delete(ctx, domain, input.Name+"."+domain, powerdns.RRTypeA)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Record wurde gelöscht",
	})
}
