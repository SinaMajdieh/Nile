package database

// Functions on selecting data from the users table in the sqlite database file

import "Nile/internal/models"

// Selecting user by user name from the database returning a slice of user model
func (db SqliteDb) SelectUserByUsername(username string) []models.User {

	// Query database
	rows, err := db.db.Query("SELECT * FROM users WHERE username = ?", &username)

	// Catching the error
	if err != nil {
		logger.Println(err)
		return nil
	}
	// Closing the rows
	defer rows.Close()

	// A user slice to hold data from returned rows.
	var users []models.User

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Hash); err != nil {
			logger.Println(err)
			return nil
		}

		// There was no errors
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		logger.Println(err)
		return nil
	}

	// There was no errors returning the value
	return users
}
