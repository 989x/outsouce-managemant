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

func (h staffHandler) ListStaffs(c *fiber.Ctx) error {
	result, err := h.staffSrv.SrvGetAllStaff()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"result":       "fail",
			"data":         nil,
			"errorMessage": err.Error(),
			"code":         400,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result":       "fial",
		"data":         result,
		"errorMessage": nil,
		"code":         200,
	})
}
