package helpers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

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

func SplitParser(input_string string) []interface{} {
	array_string := strings.Split(input_string, ",")

	interface_string := make([]interface{}, len(array_string))

	for index_array, value_array := range array_string {
		interface_string[index_array] = value_array
	}

	return interface_string
}

func SearchCodition(search []interface{}, slide string) bool {
	for searchIndex, _ := range search {
		if search[searchIndex] == "All" || search[searchIndex] == "" || search[searchIndex] == slide {
			return true
		}
	}
	return false
}
