package repository

import (
	"osm/api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StaffRepository interface {
	GetStaff() ([]models.StaffDashBoard, error)
	SetPrimetiveFilter([]interface{}, []interface{}, []interface{}, []interface{}) (primitive.D, error)
	GetPipeLinePrimetive(primitive.D, time.Time) ([]models.StaffDashBoard, error)
	GetCountCenterStaff(string, string, time.Time) ([]models.StaffCenterStatus, error)
}
