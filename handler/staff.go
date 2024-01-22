package handler

import (
	service "osm/service/staff_service"

	"github.com/gofiber/fiber/v2"
)

type staffHandler struct {
	staffSrv service.StaffService
}

func NewStaffService(staffSrv service.StaffService) staffHandler {
	return staffHandler{staffSrv: staffSrv}
}

func (h staffHandler) GetDashboard(c *fiber.Ctx) error {
	result, err := h.staffSrv.SrvGetDashboard()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"result":       "Fail",
			"data":         nil,
			"errorMessage": err.Error(),
			"code":         400,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result":       "Success",
		"data":         result,
		"errorMessage": nil,
		"code":         200,
	})
}

func (h staffHandler) ListStaffs(c *fiber.Ctx) error {
	result, err := h.staffSrv.SrvGetAllStaff()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"result":       "Fail",
			"data":         nil,
			"errorMessage": err.Error(),
			"code":         400,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result":       "Success",
		"data":         result,
		"errorMessage": nil,
		"code":         200,
	})
}

func (h staffHandler) ReadStaff(c *fiber.Ctx) error {

	paramsId := c.Params("id")
	if paramsId == "" {
		return c.Status(400).JSON(fiber.Map{
			"result":       "Fail",
			"data":         nil,
			"errorMessage": "Params id is empty",
			"code":         400,
		})
	}

	result, err := h.staffSrv.SrvGetStaffById(paramsId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"result":       "Fail",
			"data":         nil,
			"errorMessage": err.Error(),
			"code":         400,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result":       "Success",
		"data":         result,
		"errorMessage": nil,
		"code":         200,
	})
}
