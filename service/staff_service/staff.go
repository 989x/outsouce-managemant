package service

import "osm/models"

type StaffService interface {
	SrvGetDashboard() ([]models.StaffJobResponse, error)
	SrvGetAllStaff() ([]models.StaffResponse, error)
	SrvGetStaffById(string) (*models.Staff, error)
}
