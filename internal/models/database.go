// package models consists of all the models for the database interface,
// front-end forms, ...
package models

// Models used in sqlite database

// User model for sqlite users table
type User struct {
	// User's username
	Id int `sql:"id"`

	// User's
	Username string `sql:"username"`

	// User's password hash
	Hash string `sql:"hash"`
}

// Database interface for api to use
type DataBase interface {
	Open(string) error
	SelectUserByUsername(string) []User
	InsertUser(*RegisterForm) bool
}
