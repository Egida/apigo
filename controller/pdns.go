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

type Zone struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Type             string   `json:"type,omitempty"`
	URL              string   `json:"url"`
	Kind             string   `json:"kind"`
	Serial           int      `json:"serial"`
	NotifiedSerial   int      `json:"notified_serial"`
	EditedSerial     int      `json:"edited_serial"`
	Masters          []string `json:"masters"`
	DNSSEC           bool     `json:"dnssec"`
	NSEC3Param       string   `json:"nsec3param,omitempty"`
	NSEC3Narrow      bool     `json:"nsec3narrow,omitempty"`
	Presigned        bool     `json:"presigned,omitempty"`
	SOAEdit          string   `json:"soa_edit,omitempty"`
	SOAEditAPI       string   `json:"soa_edit_api,omitempty"`
	APIRectify       bool     `json:"api_rectify,omitempty"`
	Zone             string   `json:"zone,omitempty"`
	Catalog          string   `json:"catalog,omitempty"`
	Account          string   `json:"account,omitempty"`
	NameServers      []string `json:"nameservers,omitempty"`
	MasterTSIGKeyIDs []string `json:"master_tsig_key_ids,omitempty"`
	SlaveTSIGKeyIDs  []string `json:"slave_tsig_key_ids,omitempty"`
}

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
func AddZZ(c *fiber.Ctx) error {
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
	err := pwdns.Records.Delete(ctx, domain, input.Name+"."+domain, powerdns.RRType(input.Type))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Record wurde gelöscht",
	})
}
