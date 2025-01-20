package routes

import (
	"github.com/microsite-ilustrasi/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupProtectedRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/protected", middleware.Protected(), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "This is a protected route",
		})
	})
}
