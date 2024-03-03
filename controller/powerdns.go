package controller

import (
	"api/model"
	"api/pdns"
	"api/synlinq"
	"context"
	"net"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joeig/go-powerdns/v3"
	"github.com/spf13/viper"
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
func ChangePtr(c *fiber.Ctx) error {
	ip := c.Params("ip")
	ips := net.ParseIP(ip)

	head := c.GetReqHeaders()
	token := head["X-Apikey"]
	ipadress, err := model.FindByip(ip)
	isused, err := model.FindAPIKey(token)
	usedemail, err := model.FindUserById(isused.UserID)
	ips = net.ParseIP(ip)

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, ipadress.Customer)
	}
	if ipadress.Ip == ip {
		if ipadress.Customer == usedemail.Email || usedemail.IsAdmin == true {
			var input model.RecordIn
			if ips.To4() != nil {
				if err := c.BodyParser(&input); err != nil {
					return fiber.NewError(fiber.StatusBadRequest, err.Error())
				}

				if err := model.Validate.Struct(&input); err != nil {
					return err
				}
				//IPv4 Request Synlinq
				_, err := synlinq.AddPtr(ipadress.Ip, input.Data)
				if err != nil {
					return fiber.NewError(fiber.StatusBadRequest, err.Error())
				}
				return c.Status(fiber.StatusOK).JSON(fiber.Map{
					"success": true,
					"message": "Ptr was changed",
				})
			} else {
				//IPv6 Request Pdns

				if err := c.BodyParser(&input); err != nil {
					return fiber.NewError(fiber.StatusBadRequest, err.Error())
				}

				if err := model.Validate.Struct(&input); err != nil {
					return err
				}
				pwdns := powerdns.NewClient(viper.GetString("app.powerdnsserver"), "localhost", map[string]string{"X-API-Key": viper.GetString("app.powerdnskey")}, nil)
				ctx := context.Background()
				err := pwdns.Records.Change(ctx, ipadress.Zone, ipadress.Zonename+"."+ipadress.Zone, powerdns.RRTypePTR, 60, []string{input.Data + "."})
				if err != nil {
					return fiber.NewError(fiber.StatusBadRequest, err.Error())
				}

				return c.Status(fiber.StatusOK).JSON(fiber.Map{
					"success": true,
					"message": "Ptr was changed",
				})
			}

		} else {

			return fiber.NewError(fiber.StatusBadRequest, "")
		}

	} else {

		return fiber.NewError(fiber.StatusBadRequest, "Zone not found")
	}
	return nil
}
