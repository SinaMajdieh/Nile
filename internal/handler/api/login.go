package api

// Api handlers related to Login actions

import (
	"Nile/internal/messages"
	"Nile/internal/models"
	"Nile/internal/validation"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

// API: Handle user login request
func LoginHandler(c *fiber.Ctx) error {

	// Declaring the login form
	loginForm := new(models.LoginForm)

	// Parse the login form
	// Catch errors
	if err := c.BodyParser(&loginForm); nil != err {
		logger.Println(messages.LoginFormError)
		return generateStatus(c, false, messages.LoginFormError)
	}

	id, message := validation.LoginIsValid(loginForm)

	// Ensure validations was true
	if id == -1 {
		return generateStatus(c, false, message)
	}

	// Start users session via setting a cookie
	c.Cookie(&fiber.Cookie{
		Name:  "user_id",
		Value: strconv.Itoa(id),
	})

	logger.Println(loginForm.Username + " logged in")
	return generateStatus(c, true, messages.LoginSuccessful)

}
