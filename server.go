/* Main function for the WEB API server.

This function initializes and runs the web server for the WEB application.
It performs the following operations:
- Loads environment variables from .env file
- Initializes application configuration
- Connects to the Oracle database
- Sets up the Fiber web framework with CORS support
- Configures API routes for patient-related operations
- Starts the HTTP server on the configured port

Note: The following code blocks are commented out but would be used when integrating with a frontend:
- app.Static("/", cfg.AppPath) - Would serve static files from the frontend build
- app.Get("/site/*", ...) - Would handle SPA routing by serving the index.html for all routes

The server listens on all interfaces (0.0.0.0) using the port specified in the configuration.
*/

package main

import (
	"fmt"

	"log"

	"go-web-template/internal/dummy" // Update this path to match your project structure
	"go-web-template/utils"          // Update this path to match your project structure

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// read .env file with godotenv
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cfg, err := utils.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	// connect to oracle database
	db := utils.DBConnector{Config: *cfg}
	conn, err := db.Connect()
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer conn.Close()

	app := fiber.New()

	// Serve static files from the frontend build
	// app.Static("/", cfg.AppPath)

	// Use Fiber's CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                // Your React app’s origin
		AllowMethods: "GET,POST,OPTIONS", // Allowed HTTP methods
		AllowHeaders: "Content-Type",     // Allowed headers
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Connected")
	})

	// API routes
	route := app.Group("/site/api")
	dummy.Routes(&route, conn)

	// handle Single-Page Application routing
	// app.Get("/site/*", func(c *fiber.Ctx) error {
	// 	return c.SendFile(filepath.Join(cfg.AppPath, "index.html"))
	// })

	// run app server
	log.Fatal(app.Listen("0.0.0.0:" + cfg.ServerPort))

}
