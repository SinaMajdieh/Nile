package database

// Functions on inserting data to the users table in the sqlite database file

import (
	"Nile/internal/models"

	"golang.org/x/crypto/bcrypt"
)

// Hash difficulty
// Used for encrypting the user password
const(
	hashDifficulty = 6
)

// Insert a new user to the users table
// Returns true if the insertion was successful
// Returns false if it was not
func (db *SqliteDb) InsertUser(form *models.RegisterForm) bool {

	// Preparing the sqlite for the insertion query
	// catching the error
	insertSQL := `INSERT INTO users (username, hash) VALUES (?, ?)`
	statement, err := db.db.Prepare(insertSQL)
	if nil != err {
		logger.Println(err)
		return false
	}

	// Generating the hash
	// catching the error
	hashByte, err := bcrypt.GenerateFromPassword([]byte(form.Password), hashDifficulty)
	if nil != err {
		logger.Println(err)
		return false
	}
	hash := string(hashByte)

	// Query database
	// Catching the error
	_, err = statement.Exec(form.Username, hash)
	if nil != err {
		logger.Println(err)
		return false
	}

	return true
}
