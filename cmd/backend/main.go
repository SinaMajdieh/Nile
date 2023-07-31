package main

// This is the main file with the main function in it
// Preparing the database and server
// Alongside with serving the server

import (
	"Nile/database"
	"Nile/internal/config"
	fibServer "Nile/internal/server"
	"log"

	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
)

// Initialize server interface
type server interface {
	// Prepare a Go app
	// Including the standard html engine
	// And cors configuration
	// While enabling cookie encryption
	PrepareServer(string)

	// Serve static files (Directory)
	SetStaticRoute(string, string)

	// Add Routers
	SetRoutes()

	// Initializing the api
	InitApi()

	// Listen And Serve the server framework
	ListenAndServe() error
}

func main() {

	// Read the config file for necessary configurations
	// Including host name which is the ip
	// And port which is the port
	config, err := config.LoadServerConfig("./cmd/config.json")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Sqlite Database
	database.Db = database.InitSqliteDb(config)

	// Initialize Fiber Server
	var app server
	app = fibServer.InitializeServer(config)

	// Preparing the server
	// Set cooky encryption generator
	cookieGeneratorKey := encryptcookie.GenerateKey()
	app.PrepareServer(cookieGeneratorKey)

	// Initialize static files serving
	app.SetStaticRoute("/", "./build")

	// Initialize Routes
	app.SetRoutes()

	// Initialize api
	app.InitApi()

	// Serving and Listening
	log.Fatal(app.ListenAndServe())
}
