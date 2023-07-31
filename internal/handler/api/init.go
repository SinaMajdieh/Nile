// Package api is for running and handling the api request.
// designed to be used in this project alone
package api

// Initialize api basic functions

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

// Logger For api
var logger = log.New(os.Stdout, "API: ", log.Ldate|log.Ltime)

// Generating a json message for sending back to client.
// Content of the message includes status and the message text.
func generateStatus(c *fiber.Ctx, status bool, message string) error {
	return c.JSON(&fiber.Map{
		"status":  status,
		"message": message,
	})
}
