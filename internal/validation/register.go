package validation

// Validation functions for registration action

import (
	"Nile/database"
	"Nile/internal/messages"
	"Nile/internal/models"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "VALIDATION: ", log.Ldate|log.Ltime)

// Validates Register request By checking wether the username is taken
// And whether the password matches the confirmation
func RegisterIsValid(form *models.RegisterForm) (bool, string) {
	// Query database
	// Catch errors
	users := database.Db.SelectUserByUsername(form.Username)
	// if nil == users {
	// 	logger.Println("no rows")
	// }

	// Ensure username is not taken
	if len(users) != 0 {
		return false, messages.UserNameTaken
	}

	// Ensure password matches confirmation
	if form.Password != form.Confirmation {
		return false, messages.PasswordMatchProblem
	}

	return true, ""
}
