package helpers

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

// HandleError logs the error and returns a JSON response with an error message
func HandleError(c *fiber.Ctx, status int, errMsg string, err error) error {
	log.Println(err)
	return c.Status(status).JSON(fiber.Map{"error": errMsg})
}
