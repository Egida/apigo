package error

import (
	"api/model"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx, err error) error {
	status := fiber.StatusInternalServerError
	description := err.Error()

	switch e := err.(type) {
	case *fiber.Error:
		status = e.Code

	case validator.ValidationErrors:
		status = fiber.StatusBadRequest

		var stringErrors []string
		for _, err := range e {
			stringErrors = append(stringErrors, validationErrorToText(err))
		}

		description = strings.Join(stringErrors, "; ")
	}

	if strings.HasPrefix(c.Path(), "/admin") {
		return c.Status(status).Render("error", fiber.Map{
			"Status": status,
			"Error":  description,
		})
	}

	return c.Status(status).JSON(&model.Error{
		Error:            http.StatusText(status),
		ErrorCode:        status,
		ErrorDescription: description,
		{"Success":false}
	})
}

func validationErrorToText(e validator.FieldError) string {
	fieldName := strings.ToLower(e.Field())
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("field '%s' is required", fieldName)
	case "max":
		return fmt.Sprintf("field '%s' must be less or equal to %s", fieldName, e.Param())
	case "min":
		return fmt.Sprintf("field '%s' must be more or equal to %s", fieldName, e.Param())
	}
	return fmt.Sprintf("field '%s' is not valid", fieldName)
}
