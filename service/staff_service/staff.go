package service

import "osm/models"

type StaffService interface {
	SrvGetAllStaff() ([]models.StaffResponse, error)
}
