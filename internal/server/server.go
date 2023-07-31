// Package fibServer is an implementation of a Standard Go Fiber framework
// For preparing and serving a fiber application
// designed to run a fiber server for this project
package fibServer

// Package for initializing Standard Go Fiber server app

import (
	"Nile/internal/config"
	"Nile/internal/handler/api"
	"Nile/internal/handler/pages"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/template/html/v2"
)

// Standard Go fiber server model
type FibServer struct {
	// Value server will be hosted on
	host string

	// Port value for the FabServer
	port string

	// Client Port
	clientPort string

	// Standard Go Fiber App
	app *fiber.App
}

// Initialize New FabServer
// Using the given configuration
func InitializeServer(config *config.Config) *FibServer {
	return &FibServer{
		host:       config.Host,
		port:       config.Port,
		clientPort: config.ClientPort,
		app:        nil,
	}
}

// Prepare a Go fiber app
// Including the standard Go html engine
// And cors configuration
// While enabling cookie encryption
func (server *FibServer) PrepareServer(cookieGenerator string) {
	// Initialize standard Go html template engine
	engine := html.New("./build", ".html")

	// Initialize Go fiber app for serving the api, statics, ...
	// Include the standard Go html engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Set cors configuration for the server
	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOriginsFunc: nil,
		AllowOrigins:     "http://" + server.host + server.clientPort,
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders:     "Origin, Content-Type, Accept, Accept-Language, Content-Length",
		AllowCredentials: true,
		ExposeHeaders:    "",
	}))

	// Set standard Cooky configuration
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: cookieGenerator,
	}))

	// Set app
	server.app = app
}

// Go Fiber serve static files (Directory)
func (server *FibServer) SetStaticRoute(url, path string) {
	server.app.Static(url, path)
}

// Go Fiber Add Routers
func (server *FibServer) SetRoutes() {
	// Initialize index handler
	server.app.Get("/login", pages.ReactDefaultHandler)
}

// Go Fiber Add Api
func (server *FibServer) InitApi() {
	server.app.Post("/api/login", api.LoginHandler)
	server.app.Post("/api/register", api.RegisterHandler)

	server.app.Post("/api/add-to-cart", api.AddToCart)
	server.app.Get("/api/cart", api.GetMyCart)
	server.app.Delete("/api/drop-cart", api.EmptyCart)
}

// Listen And Serve the Go Fiber framework
func (server FibServer) ListenAndServe() error {
	return server.app.Listen(server.host + server.port)
}
