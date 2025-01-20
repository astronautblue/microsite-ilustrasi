package handlers

import (
	"fmt"
	"time"

	"github.com/microsite-ilustrasi/config"
	"github.com/microsite-ilustrasi/database"
	"github.com/microsite-ilustrasi/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type LoginRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

func Login(c *fiber.Ctx) error {
	var request LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"responseCode":      "400",
			"responseMessage":   "Invalid request body",
			"responseExecution": "Failed",
		})
	}

	var user models.User
	result := database.DB.Where("username = ?", request.Username).First(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"responseCode":      "401",
			"responseMessage":   "Invalid credentials",
			"responseExecution": "Failed",
		})
	}

	if request.Password != user.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"responseCode":      "401",
			"responseMessage":   "Invalid credentials",
			"responseExecution": "Failed",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	fmt.Println("result", claims)
	t, err := token.SignedString([]byte(config.AppConfig.JWT_SECRET))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"responseCode":      "500",
			"responseMessage":   "Could not generate token",
			"responseExecution": "Failed",
		})
	}

	return c.JSON(fiber.Map{
		"token": t,
		"responseData": fiber.Map{
			"id":         user.ID,
			"username":   user.Username,
			"created_at": user.CreatedAt,
		},
		"responseCode":      "200",
		"responseMessage":   "Login successful",
		"responseExecution": "Success",
	})
}
