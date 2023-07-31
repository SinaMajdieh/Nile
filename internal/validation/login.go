// Package validation is used to validate information from the front-end.
// Used in api package
package validation

// Validation functions for login action

import (
	"Nile/database"
	"Nile/internal/messages"
	"Nile/internal/models"

	"golang.org/x/crypto/bcrypt"
)

// Ensures Login is valid via checking wether the user exist and if the password matches.
// Returns user id if Login was valid Return -1 if it was not also returning a message of the corresponding error, empty string is there was no errors
func LoginIsValid(form *models.LoginForm) (int, string) {

	// There was no errors
	// Query database for username
	// Catching the error
	users := database.Db.SelectUserByUsername(form.Username)
	if nil == users {
		logger.Println(messages.EmptyRows)
		return -1, messages.EmptyRows
	}

	// Ensure user was found
	if len(users) != 1 {
		logger.Println(messages.UserNotFound)
		return -1, messages.UserNotFound
	}

	// Ensure Password matches hash
	err := bcrypt.CompareHashAndPassword([]byte(users[0].Hash), []byte(form.Password))
	if nil != err {
		return -1, messages.WrongPassword
	}

	return users[0].Id, ""
}
