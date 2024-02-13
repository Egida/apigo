package controller

import (
	"api/active"

	"api/model"
	"api/strukt"

	"log"

	"github.com/gofiber/fiber/v2"
)

func GetPathIncidents(c *fiber.Ctx) error {
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
			routing, err := active.GetIncidents(ip)
			if err != nil {
				log.Println("ddos.Test error: ", err)
				log.Printf("ddos.Test error: %v %T\n", err, err)
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data":    routing.Data,
			})
		} else {
			return fiber.NewError(fiber.StatusForbidden, "You not authorized using this IP")

		}
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "IP not in our System")

	}
}
func GetPathRules(c *fiber.Ctx) error {
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
			routing, err := active.GetRules(ip)
			if err != nil {
				log.Println("ddos.Test error: ", err)
				log.Printf("ddos.Test error: %v %T\n", err, err)
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data":    routing,
			})
		} else {
			return fiber.NewError(fiber.StatusForbidden, "You not authorized using this IP")

		}
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "IP not in our System")

	}
}
func AddPathRules(c *fiber.Ctx) error {
	var input strukt.AddRule
	ippref := c.Params("ip")
	ipadress, err := model.FindByip(ippref)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
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
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := model.Validate.Struct(&input); err != nil {
		return err
	}

	if _, err := active.AddRule(input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Rule was added successfully",
	})
}
func GetRateLimits(c *fiber.Ctx) error {
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
			routing, err := active.GetRateLimiters()
			if err != nil {
				log.Println("ddos.Test error: ", err)
				log.Printf("ddos.Test error: %v %T\n", err, err)
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data":    routing.RateLimiters,
			})
		} else {
			return fiber.NewError(fiber.StatusForbidden, "You not authorized using this IP")

		}
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "IP not in our System")

	}
}
func DeleteRule(c *fiber.Ctx) error {
	ip := c.Params("ip")
	id := c.Params("id")
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
			routing, err := active.DeleteRule(id)
			if err != nil {
				log.Println("ddos.Test error: ", err)
				log.Printf("ddos.Test error: %v %T\n", err, err)
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data":    routing,
			})
		} else {
			return fiber.NewError(fiber.StatusForbidden, "You not authorized using this IP")

		}
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "IP not in our System")

	}
}
func GetFilters(c *fiber.Ctx) error {
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
			routing, err := active.GetFilter()
			if err != nil {
				log.Println("ddos.Test error: ", err)
				log.Printf("ddos.Test error: %v %T\n", err, err)
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data":    routing,
			})
		} else {
			return fiber.NewError(fiber.StatusForbidden, "You not authorized using this IP")

		}
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "IP not in our System")

	}
}
func AvailableFilter(c *fiber.Ctx) error {
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
			routing, err := active.AvailableFilters()
			if err != nil {
				log.Println("ddos.Test error: ", err)
				log.Printf("ddos.Test error: %v %T\n", err, err)
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data":    routing,
			})
		} else {
			return fiber.NewError(fiber.StatusForbidden, "You not authorized using this IP")

		}
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "IP not in our System")

	}
}
func AddFilter(c *fiber.Ctx) error {
	var input strukt.AddFilter
	ippref := c.Params("ip")
	ft := c.Params("filter_type")
	ipadress, err := model.FindByip(ippref)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
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
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := model.Validate.Struct(&input); err != nil {
		return err
	}

	if _, err := active.AddFilter(input, ft); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Rule was added successfully",
	})
}
func DeleteFilter(c *fiber.Ctx) error {
	ip := c.Params("ip")
	id := c.Params("id")
	ft := c.Params("filter_type")
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
			routing, err := active.DeleteFilter(id, ft)
			if err != nil {
				log.Println("ddos.Test error: ", err)
				log.Printf("ddos.Test error: %v %T\n", err, err)
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data":    routing,
			})
		} else {
			return fiber.NewError(fiber.StatusForbidden, "You not authorized using this IP")

		}
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "IP not in our System")

	}
}
func AddRateLimit(c *fiber.Ctx) error {
	var input strukt.CreateRatelimit
	ippref := c.Params("ip")
	ipadress, err := model.FindByip(ippref)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
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
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := model.Validate.Struct(&input); err != nil {
		return err
	}

	if _, err := active.AddRatelimit(input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Rule was added successfully",
	})
}
func DeleterRateLimit(c *fiber.Ctx) error {
	ip := c.Params("ip")
	id := c.Params("id")
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
			routing, err := active.DeleteRatelimit(id)
			if err != nil {
				log.Println("ddos.Test error: ", err)
				log.Printf("ddos.Test error: %v %T\n", err, err)
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data":    routing,
			})
		} else {
			return fiber.NewError(fiber.StatusForbidden, "You not authorized using this IP")

		}
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "IP not in our System")

	}
}
