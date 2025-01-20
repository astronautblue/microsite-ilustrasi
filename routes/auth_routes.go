package routes

import (
	"github.com/microsite-ilustrasi/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	auth := app.Group("/api/auth")
	auth.Post("/login", handlers.Login)
}
