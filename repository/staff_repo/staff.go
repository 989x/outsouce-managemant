package repository

import (
	"bytes"
	"osm/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StaffRepository interface {
	GetDashBoard() ([]models.StaffJobResponse, error)
	Create(*models.Staff) (*mongo.InsertOneResult, error)
	CallUserExists(buffer *bytes.Buffer) ([]models.StaffExists, error)
	GetAll() ([]models.StaffResponse, error)
	GetById(primitive.ObjectID) (*models.Staff, error)
	GetAllStaffJobs() ([]models.StaffJobsFind, error)
	Update(primitive.ObjectID) error
	Remove(primitive.ObjectID) error
}
