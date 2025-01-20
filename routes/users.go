package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/microsite-ilustrasi/database"
	"github.com/microsite-ilustrasi/models"
)

type reqBody struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

func userListHandler(c *fiber.Ctx) error {
	var request reqBody

	// Parse the incoming request body into the 'request' struct
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"responseCode":      "400",
			"responseMessage":   "Invalid request body",
			"responseExecution": "Failed",
		})
	}

	// Declare a user object to hold the user data from the database
	var user models.User

	// Query the database to find the user with the given username
	result := database.DB.Where("username = ?", request.Username).First(&user)

	// Check if there is an error in finding the user
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"responseCode":      "404",
			"responseMessage":   "User not found",
			"responseExecution": "Failed",
		})
	}

	// Return the user data in the response
	return c.JSON(fiber.Map{
		"responseData": fiber.Map{
			"id":         user.ID,
			"username":   user.Username,
			"created_at": user.CreatedAt,
		},
		"responseCode":      "200",
		"responseMessage":   "Successful",
		"responseExecution": "Success",
	})
}
