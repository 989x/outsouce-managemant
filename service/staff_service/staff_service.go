package service

import (
	"errors"
	"osm/models"
	repository "osm/repository/staff_repo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type staffService struct {
	staffRepo repository.StaffRepository
}

func NewStaffService(staffRepo repository.StaffRepository) StaffService {
	return staffService{staffRepo: staffRepo}
}

func (s staffService) SrvGetDashboard() ([]models.StaffJobResponse, error) {
	dashboard_result, err := s.staffRepo.GetDashBoard()
	if err != nil {
		return nil, err
	}
	return dashboard_result, nil
}

func (s staffService) SrvGetAllStaff() ([]models.StaffResponse, error) {
	staffs_result, err := s.staffRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return staffs_result, nil
}

func (s staffService) SrvGetStaffById(params string) (*models.Staff, error) {
	_id, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		return nil, errors.New("Inavalid object params")
	}

	staff_result, err := s.staffRepo.GetById(_id)
	if err != nil {
		return nil, errors.New("Can't find staff")
	}
	return staff_result, nil
}
