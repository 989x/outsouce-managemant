package service

import (
	"osm/api/models"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	SrvGetAllUser(*fiber.Ctx) ([]models.UserResponse, error)
	SrvGetById(*fiber.Ctx, string) ([]models.UserResponse, error)
}
