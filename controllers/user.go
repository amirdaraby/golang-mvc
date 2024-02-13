package controllers

import (
	"fmt"
	"log"

	"github.com/amirdaraby/golang-mvc/database"
	"github.com/amirdaraby/golang-mvc/models"
	"github.com/amirdaraby/golang-mvc/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type XValidator struct {
	validator *validator.Validate
}

var validate = validator.New()

func (v XValidator) Validate(data any) []response.ErrorResponse {
	validationErrors := []response.ErrorResponse{}

	errs := validate.Struct(data)

	if errs != nil {

		for _, err := range errs.(validator.ValidationErrors) {
			var elem response.ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)

		}

	}

	return validationErrors
}

func StoreUser(c *fiber.Ctx) error {

	user := new(models.User)

	err := c.BodyParser(user)

	if err != nil {

		log.Printf("user creation failed %v", err)

		return c.Status(fiber.StatusBadRequest).JSON(response.Json{
			Message: "BAD REQUEST",
		})
	}

	xValidator := &XValidator{
		validator: validate,
	}

	if errs := xValidator.Validate(user); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag))
		}

		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.Json{
			Data:    errMsgs,
			Message: "unproccessable entity dadi",
		})
	}

	err = database.DbConnection.Create(user).Error

	if err != nil {

		log.Printf("user creation failed %v", err)

		return c.Status(fiber.StatusBadRequest).JSON(response.Json{
			Message: "BAD REQUEST",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response.Json{
		Data:    *user,
		Message: "user created",
	})

}
