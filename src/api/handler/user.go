package handler

import (
	service "osm/api/service/user_service"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type userHandler struct {
	// userSrv service.UserService
	userSrv service.UserService
}

func NewUserHandler(userSrv service.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}

func (h userHandler) GetAll(c *fiber.Ctx) error {
	result, err := h.userSrv.SrvGetAllUser(c)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"result":       "Fail",
			"errorMessage": err.Error(),
			"data":         nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result":       "Success",
		"errorMessage": nil,
		"data":         result,
	})

}

func (h userHandler) GetById(c *fiber.Ctx) error {

	result, err := h.userSrv.SrvGetById(c, c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"result":       "Fail",
			"errorMessage": err.Error(),
			"data":         nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result":       "Success",
		"errorMessage": nil,
		"data":         result,
	})
}

func (h userHandler) ParamsUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return c.Status(200).JSON(fiber.Map{
		"code":   200,
		"data":   claims,
		"result": "Success",
	})
}
