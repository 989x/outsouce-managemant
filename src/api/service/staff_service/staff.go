package service

import "osm/api/models"

type StaffService interface {
	SrvGetDashboard() ([]models.StaffDashBoard, error)
	SrvGetAllStaff() ([]models.StaffResponse, error)
	SrvGetStaffById(string) (*models.Staff, error)
}
