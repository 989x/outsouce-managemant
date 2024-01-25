package helpers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// HandleError logs the error and returns a JSON response with an error message
func HandleError(c *fiber.Ctx, status int, errMsg string, err error) error {
	log.Println(err)
	return c.Status(status).JSON(fiber.Map{"error": errMsg})
}
