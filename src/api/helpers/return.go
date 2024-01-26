package helpers

import "github.com/gofiber/fiber/v2"

func JsonResponse(c *fiber.Ctx, errorMessage error, code int, data interface{}, status string) error {
	if errorMessage != nil {
		return c.Status(code).JSON(fiber.Map{
			"result":       status,
			"data":         nil,
			"errorMessage": errorMessage.Error(),
			"code":         code,
		})
	}
	return c.Status(code).JSON(fiber.Map{
		"result":       status,
		"data":         data,
		"errorMessage": nil,
		"code":         code,
	})
}
