package controller

import (
	"api/combahton"
	"api/model"
	"api/strukt"
	"fmt"
	"io"

	//"io/ioutil"
	"mime/multipart"
	"time"

	"log"

	"github.com/gofiber/fiber/v2"
)

func GetRouting(c *fiber.Ctx) error {
	ip := c.Params("ip")
	ipadress, err := model.FindByip(ip)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, ipadress.Customer)
	}
	head := c.GetReqHeaders()
	token := head["X-Apikey"]

	isused, err := model.FindAPIKey(token)

	usedemail, err := model.FindUserById(isused.UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "User not found")
	}

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if ipadress.Ip == ip {

		if ipadress.Customer == usedemail.Email || usedemail.IsAdmin == true {
			routing, err := combahton.GetRouting(ip)
			if err != nil {
				log.Println("ddos.Test error: ", err)
				log.Printf("ddos.Test error: %v %T\n", err, err)
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"status": fiber.StatusOK,
				"result": routing,
			})
		} else {
			return fiber.NewError(fiber.StatusForbidden, "You not authorized using this IP")

		}
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "IP not in our System")

	}
}
func AddRouting(c *fiber.Ctx) error {
	var input strukt.DDOSLayer4
	ippref := c.Params("ip")
	mask := c.Params("mask")
	ipadress, err := model.FindByip(ippref)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, ipadress.Customer)
	}
	head := c.GetReqHeaders()
	token := head["X-Apikey"]

	isused, err := model.FindAPIKey(token)

	usedemail, err := model.FindUserById(isused.UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "User not found")
	}

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if ipadress.Ip != ippref {
		return fiber.NewError(fiber.StatusBadRequest, "IP not in our System")
	}

	if ipadress.Customer != usedemail.Email && !usedemail.IsAdmin {
		return fiber.NewError(fiber.StatusForbidden, "You not authorized using this IP")
	}

	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, ippref)
	}

	input.Prefix = fmt.Sprintf("%s/%s", ippref, mask)
	if err := model.Validate.Struct(&input); err != nil {
		return err
	}

	if _, err := combahton.AddRouting(input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, ippref)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Route wurde geändert",
	})
}
func AddVhost(c *fiber.Ctx) error {
	var input model.Layer7Input
	ippref := c.Params("ip")
	mask := c.Params("mask")
	ipadress, err := model.FindByip(ippref)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, ipadress.Customer)
	}
	head := c.GetReqHeaders()
	token := head["X-Apikey"]

	isused, err := model.FindAPIKey(token)

	usedemail, err := model.FindUserById(isused.UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "User not found")
	}

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if ipadress.Ip != ippref {
		return fiber.NewError(fiber.StatusBadRequest, "IP not in our System")
	}

	if ipadress.Customer != usedemail.Email && !usedemail.IsAdmin {
		return fiber.NewError(fiber.StatusForbidden, "You not authorized using this IP")
	}

	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, ippref)
	}

	input.IPAddress = fmt.Sprintf("%s/%s", ippref, mask)
	if err := model.Validate.Struct(&input); err != nil {
		return err
	}
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := model.Validate.Struct(&input); err != nil {
		return err
	}

	if input.Certificate == "" {
		certFile, err := c.FormFile("certificate")
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		cert, err := readMultipartFile(certFile)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		input.Certificate = cert
	}

	if input.PrivateKey == "" {
		keyFile, err := c.FormFile("privatekey")
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		pkey, err := readMultipartFile(keyFile)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		input.PrivateKey = pkey
	}

	certificate, err := combahton.CreateCertificate(
		input.IPAddress,
		input.Domain,
		input.Certificate,
		input.PrivateKey,
		0)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// We need to wait to propagate certificate on Combahton end
	time.Sleep(1 * time.Second)

	vhost, err := combahton.CreateVhost(
		certificate.IPAddress,
		certificate.Domain,
		certificate.UUID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "Vhost wurde hinzugefügt",
		"VHID":    vhost.UUID,
	})
}

func readMultipartFile(header *multipart.FileHeader) (string, error) {
	file, err := header.Open()
	if err != nil {
		return "", err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
func GetIncidents(c *fiber.Ctx) error {
	ip := c.Params("ip")
	ipadress, err := model.FindByip(ip)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, ipadress.Customer)
	}
	head := c.GetReqHeaders()
	token := head["X-Apikey"]

	isused, err := model.FindAPIKey(token)

	usedemail, err := model.FindUserById(isused.UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "User not found")
	}

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if ipadress.Ip == ip {

		if ipadress.Customer == usedemail.Email || usedemail.IsAdmin == true {
			routing, err := combahton.GetIncidents(ip)
			if err != nil {
				log.Println("ddos.Test error: ", err)
				log.Printf("ddos.Test error: %v %T\n", err, err)
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data": routing,
			})
		} else {
			return fiber.NewError(fiber.StatusForbidden, "You not authorized using this IP")

		}
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "IP not in our System")

	}
}
