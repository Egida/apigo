package controller

import (
	"api/dynadot"
	"api/model"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Whois(c *fiber.Ctx) error {
	domain := c.Params("domain")

	output, err := dynadot.Search(domain)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	log.Info().Str("client", "dynadot").Msgf("%v+\n", output)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"domain":    output.SearchResponse.SearchResults[0].DomainName,
		"available": output.SearchResponse.SearchResults[0].Available,
	})
}

func AddContact(c *fiber.Ctx) error {
	var input model.ContactInput
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := model.Validate.Struct(&input); err != nil {
		return err
	}

	output, err := dynadot.CreateContact(input)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	log.Info().Str("client", "dynadot").Msgf("%v+\n", output)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":    output.CreateContactResponse.Status,
		"contactid": output.CreateContactResponse.CreateContactContent.ContactId,
	})
}
