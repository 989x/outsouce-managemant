package service

import (
	"osm/api/models"
	"time"
)

type StaffService interface {
	SrvGetSatffDashboard(time.Time) (*models.CountStaffDashBoard, error)
	SrvGetStaff() ([]models.StaffDashBoard, error)
}
