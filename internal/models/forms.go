package models

// Forms used to transfer data from front-end to api/back-end

// Register Form model
type RegisterForm struct {
	// Username of the user
	Username     string `json:"username"`
	
	// Password the user put in
	Password     string `json:"password"`
	
	// Password confirmation to make sure user didn't make a typo
	Confirmation string `json:"confirmation"`
}

// Login Form model
type LoginForm struct {
	// Username of the user
	Username string `json:"username"`
	
	// User's password for login validation
	Password string `json:"password"`
}
