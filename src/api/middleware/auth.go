package middleware

import (
	"net/http"

	jwt "github.com/form3tech-oss/jwt-go"

	"github.com/gofiber/fiber/v2"
)

// RequireAdmin Ensures A route Can Only Be Accessed by an Admin user
// This function can be extended to handle different roles
func RequireAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	var errorList []*fiber.Error

	if role != "admin" {
		errorList = append(
			errorList,
			&fiber.Error{
				Code:    fiber.StatusUnauthorized,
				Message: "You're Not Authorized",
			},
		)
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"errors": errorList})
	}
	return c.Next()
}
