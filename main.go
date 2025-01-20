package main

import (
	"fmt"
	"log"
	"os"

	"github.com/microsite-ilustrasi/database"
	"github.com/microsite-ilustrasi/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))

	// Initialize database
	database.InitDatabase()

	// Create Fiber app
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Setup routes
	routes.SetupAuthRoutes(app)
	routes.SetupProtectedRoutes(app)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
