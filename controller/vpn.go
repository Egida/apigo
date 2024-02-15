package controller

import (
	"api/vpn"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func GetAccounts(c *fiber.Ctx) error {

	output, err := vpn.GetAccounts()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	log.Info().Str("client", "vpn").Msgf("%v+\n", output)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		&output.Data,
	})
}
