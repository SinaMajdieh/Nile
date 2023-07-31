package api

// Api handlers related to registration actions

import (
	"Nile/database"
	"Nile/internal/messages"
	"Nile/internal/models"
	"Nile/internal/validation"
	"log"

	"github.com/gofiber/fiber/v2"
)

// API: Handle user register request
func RegisterHandler(c *fiber.Ctx) error {

	// Declaring the register form to parse to the request
	registerForm := new(models.RegisterForm)

	// Parse the register form
	// Catch the error
	if err := c.BodyParser(&registerForm); nil != err {
		logger.Println(messages.RegisterFormError)
		return generateStatus(c, false, messages.RegisterFormError)
	}

	// There was no error
	// Ensure registration is valid
	validation, message := validation.RegisterIsValid(registerForm)
	if !validation {
		logger.Println(message)
		return generateStatus(c, false, message)
	}

	// Validation was successful
	// Query database to add the new user
	isInserted := database.Db.InsertUser(registerForm)

	// Ensuring insertion was successful
	if !isInserted {
		log.Println(messages.InsertionProblem)
		return generateStatus(c, false, messages.InsertionProblem)
	}

	return generateStatus(c, true, messages.RegistrationSuccessful)
}
