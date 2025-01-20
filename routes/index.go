package routes

import (
	"github.com/microsite-ilustrasi/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v2")
	api.Post("/users", middleware.Protected(), userListHandler)
}
