package handler

import (
	service "osm/api/service/staff_service"
	"time"

	"github.com/gofiber/fiber/v2"
)

type staffHandler struct {
	staffSrv service.StaffService
}

func NewStaffService(staffSrv service.StaffService) staffHandler {
	return staffHandler{staffSrv: staffSrv}
}

func (h staffHandler) GetStaffDashboard(c *fiber.Ctx) error {
	date_now := time.Now().Add(-24 * time.Hour)

	result, err := h.staffSrv.SrvGetSatffDashboard(date_now)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":         400,
			"result":       "Fail",
			"data":         nil,
			"errorMessage": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"code":         200,
		"result":       "Success",
		"data":         result,
		"errorMessage": nil,
	})
}

func (h staffHandler) GetStaff(c *fiber.Ctx) error {
	result, err := h.staffSrv.SrvGetStaff()
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
