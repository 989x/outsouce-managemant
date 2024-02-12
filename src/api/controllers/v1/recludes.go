package controllers

import (
	"outsource-management/api/helpers"

	"github.com/gofiber/fiber/v2"
)

func MethodGet(c *fiber.Ctx) error {
	return helpers.JsonResponse(c, nil, 200, "Hello, RecludeMethod GET endpoint", "success")
}

func MethodPost(c *fiber.Ctx) error {
	return helpers.JsonResponse(c, nil, 200, "Hello, RecludeMethod POST endpoint", "success")
}

func MethodPut(c *fiber.Ctx) error {
	return helpers.JsonResponse(c, nil, 200, "Hello, RecludeMethod PUT endpoint", "success")
}

func MethodDelete(c *fiber.Ctx) error {
	return helpers.JsonResponse(c, nil, 200, "Hello, RecludeMethod DELETE endpoint", "success")
}
