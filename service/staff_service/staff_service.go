package service

import (
	"osm/models"
	repository "osm/repository/staff_repo"
)

type staffService struct {
	staffRepo repository.StaffRepository
}

func NewStaffService(staffRepo repository.StaffRepository) StaffService {
	return staffService{staffRepo: staffRepo}
}

func (s staffService) SrvGetAllStaff() ([]models.StaffResponse, error) {
	staffs_result, err := s.staffRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return staffs_result, nil
}
