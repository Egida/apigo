package controller

import (
	"api/cloudns"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func AddCloudzone(c *fiber.Ctx) error {
	domain := c.FormValue("domain")

	output, err := cloudns.AddZone(domain)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	log.Info().Str("client", "cloudns").Msgf("%v+\n", output)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": &output,
	})
}
func DeleteCloudzone(c *fiber.Ctx) error {
	domain := c.FormValue("domain")

	output, err := cloudns.DeleteCloudZone(domain)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	log.Info().Str("client", "cloudns").Msgf("%v+\n", output)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": &output,
	})
}
func AddCloudrecord(c *fiber.Ctx) error {
	domain := c.FormValue("domain")
	r_type := c.FormValue("rtype")
	host := c.FormValue("host")
	ttl := 3600
	record := c.FormValue("record")
	prio := 10
	output, err := cloudns.AddCloudrecord(domain, r_type, host, ttl, record, prio)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	log.Error().Str("client", "cloudns").Msgf("%v+\n", output)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": &output,
	})
}
