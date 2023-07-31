// Package messages includes all the error and success messages
// used in this project only
package messages

// Declaring messages different packages will use

// errors declared for login api
const (
	LoginFormError  = "Something happened on parsing the login form ..."
	EmptyRows       = "empty rows"
	UserNotFound    = "User was not found"
	UserFound       = "User was found"
	WrongPassword   = "Wrong Username or Password"
	LoginSuccessful = "Logging in"
)

// errors declared for registration api
const (
	RegisterFormError      = "Something happened on parsing the registration form ..."
	RegisterFormNotValid   = "Registration form was not valid"
	InsertionProblem       = "Registration was not successful due to insertion problems..."
	RegistrationSuccessful = "Registration was Successful"
	UserNameTaken          = "username was taken"
	PasswordMatchProblem   = "password does not match the confirmation"
)
