package service

import (
	"errors"
	"fmt"
	"osm/api/models"
	repository "osm/api/repository/staff_repo"
	"time"
)

type staffService struct {
	staffRepo repository.StaffRepository
}

func NewStaffService(staffRepo repository.StaffRepository) StaffService {
	return staffService{staffRepo: staffRepo}
}

func (s staffService) SrvGetSatffDashboard(date_nPlus time.Time) (*models.CountStaffDashBoard, error) {

	availableValues := []interface{}{"Available"}
	teamValues := []interface{}{"Dev", "Tester", "UXUI", "Data Sci", "DevOps"}
	statusValues := []interface{}{"ลา", "ลาไม่รับเงินเดือน", "ลาคลอด", "ลาอุปสมบท", "ลารับราชการทหาร", "ลาฌาปนกิจ"}
	centerValues := []interface{}{"ขอนแก่น", "เชียงใหม่", "หาดใหญ่", "กรุงเทพ"}
	// ____________________________________________Query - Dev Available____________________________________________//
	filter_devAvanonit, err := s.staffRepo.SetPrimetiveFilter(availableValues, teamValues, statusValues, centerValues)
	if err != nil {
		return nil, errors.New("Get Dev avaliable non init")
	}
	getDevAvanonit, err := s.staffRepo.GetPipeLinePrimetive(filter_devAvanonit, date_nPlus)
	if err != nil {
		return nil, errors.New("Get Dev avaliable non init")
	}
	getDevAvanonit_count := len(getDevAvanonit)
	// ____________________________________________Query - Dev Available____________________________________________//

	availableValues = []interface{}{"On Board"}
	filter_devOnnonitt, err := s.staffRepo.SetPrimetiveFilter(availableValues, teamValues, statusValues, centerValues)
	//  ____________________________________________Query - Dev On board ____________________________________________//

	getDevOnnonit, err := s.staffRepo.GetPipeLinePrimetive(filter_devOnnonitt, date_nPlus)
	if err != nil {
		return nil, errors.New("Get Dev avaliable non init")
	}
	getDevOnnonit_count := len(getDevOnnonit)
	//  ____________________________________________Query - Dev On board ____________________________________________//

	availableValues = []interface{}{"Available"}
	teamValues = []interface{}{"IT Infra"}
	// ____________________________________________Query - It Available____________________________________________//

	filter_itAva, err := s.staffRepo.SetPrimetiveFilter(availableValues, teamValues, statusValues, centerValues)

	getItAva, err := s.staffRepo.GetPipeLinePrimetive(filter_itAva, date_nPlus)
	if err != nil {
		return nil, errors.New("Get Dev avaliable non init")
	}
	getItAva_count := len(getItAva)
	// ____________________________________________Query - It Available____________________________________________//

	availableValues = []interface{}{"On Board"}
	teamValues = []interface{}{"IT Infra"}
	// ____________________________________________Query - On Board____________________________________________//

	filter_itOn, err := s.staffRepo.SetPrimetiveFilter(availableValues, teamValues, statusValues, centerValues)

	getItOn, err := s.staffRepo.GetPipeLinePrimetive(filter_itOn, date_nPlus)
	if err != nil {
		return nil, errors.New("Get Dev avaliable non init")
	}
	getItOn_count := len(getItOn)
	// ____________________________________________Query - On Board____________________________________________//

	// ____________________________________________Site Query - BnkAva____________________________________________//

	bnkAva_slide, err := s.staffRepo.GetCountCenterStaff("กรุงเทพ", "Available", date_nPlus)
	if err != nil {
		fmt.Println(err.Error())
	}
	bnkAva_count := len(bnkAva_slide)

	fmt.Println("กรุงเทพ :Available")
	fmt.Println(bnkAva_count)
	fmt.Println(bnkAva_slide)
	// ____________________________________________Site Query - BnkAva____________________________________________//

	// ____________________________________________Site Query - BnkOnb____________________________________________//

	bnkOnb_slide, err := s.staffRepo.GetCountCenterStaff("กรุงเทพ", "On Board", date_nPlus)
	if err != nil {
		fmt.Println(err.Error())
	}
	bnkOnb_count := len(bnkOnb_slide)
	fmt.Println("กรุงเทพ : On Board")
	fmt.Println(bnkOnb_count)
	fmt.Println(bnkOnb_slide)
	// ____________________________________________Site Query - BnkOnb____________________________________________//

	// ____________________________________________Site Query - ChmAva____________________________________________//
	chmAva_slide, err := s.staffRepo.GetCountCenterStaff("เชียงใหม่", "Available", date_nPlus)
	if err != nil {
		fmt.Println(err.Error())
	}
	chmAva_count := len(chmAva_slide)
	fmt.Println("เชียงใหม่ : Available")
	fmt.Println(chmAva_count)
	fmt.Println(chmAva_slide)
	// ____________________________________________Site Query - ChmAva____________________________________________//

	// ____________________________________________Site Query - ChmOnb____________________________________________//
	chmOnb_slide, err := s.staffRepo.GetCountCenterStaff("เชียงใหม่", "On Board", date_nPlus)
	if err != nil {
		fmt.Println(err.Error())
	}
	chmOnb_count := len(chmOnb_slide)
	fmt.Println("เชียงใหม่ : On Board")
	fmt.Println(chmOnb_count)
	fmt.Println(chmOnb_slide)
	// ____________________________________________Site Query - ChmOnb____________________________________________//

	// ____________________________________________Site Query - KhnAva____________________________________________//
	khnAva_slide, err := s.staffRepo.GetCountCenterStaff("ขอนแก่น", "Available", date_nPlus)
	if err != nil {
		fmt.Println(err.Error())
	}
	khnAva_count := len(khnAva_slide)
	fmt.Println("ขอนแก่น : Available")
	fmt.Println(khnAva_count)
	fmt.Println(khnAva_slide)
	// ____________________________________________Site Query - KhnAva____________________________________________//

	// ____________________________________________Site Query - KhnOnb____________________________________________//
	khnOnb_slide, err := s.staffRepo.GetCountCenterStaff("ขอนแก่น", "On Board", date_nPlus)
	if err != nil {
		fmt.Println(err.Error())
	}
	khnOnb_count := len(khnOnb_slide)
	fmt.Println("ขอนแก่น : On Board")
	fmt.Println(khnOnb_count)
	fmt.Println(khnOnb_slide)
	// ____________________________________________Site Query - KhnOnb____________________________________________//

	// ____________________________________________Site Query - HdyAva____________________________________________//
	hdyAva_slide, err := s.staffRepo.GetCountCenterStaff("หาดใหญ่", "Available", date_nPlus)
	if err != nil {
		fmt.Println(err.Error())
	}
	hdyAva_count := len(hdyAva_slide)
	fmt.Println("หาดใหญ่ : Available")
	fmt.Println(hdyAva_count)
	fmt.Println(hdyAva_slide)
	// ____________________________________________Site Query - HdyAva____________________________________________//

	// ____________________________________________Site Query - HdyOnb____________________________________________//
	hdyOnb_slide, err := s.staffRepo.GetCountCenterStaff("หาดใหญ่", "On Board", date_nPlus)
	if err != nil {
		fmt.Println(err.Error())
	}
	hdyOnb_count := len(hdyOnb_slide)
	fmt.Println("หาดใหญ่ : On Board")
	fmt.Println(hdyOnb_count)
	fmt.Println(hdyOnb_slide)
	// ____________________________________________Site Query - HdyOnb____________________________________________//

	staff_total := models.StaffTotal{
		All:          getDevAvanonit_count + getDevOnnonit_count + getItAva_count + getItOn_count,
		AllAvailable: getDevAvanonit_count + getItAva_count,
		AllOnBoard:   getDevOnnonit_count + getItOn_count,
		DevOnBoard:   getDevOnnonit_count,
		DevAvailable: getDevAvanonit_count,
		AllDev:       getDevAvanonit_count + getDevOnnonit_count,
		ItOnBoard:    getItOn_count,
		ItAvailable:  getItAva_count,
		AllIt:        getItAva_count + getItOn_count,
	}

	staff_center := models.StaffCountCenter{
		BnkAvaSlide:  bnkAva_slide,
		BnkOnbSlide:  bnkOnb_slide,
		BnkAvaSCount: bnkAva_count,
		BnkOnbSCount: bnkOnb_count,
		ChmAvaSlide:  chmAva_slide,
		ChmOnbSlide:  chmOnb_slide,
		ChmAvaSCount: chmAva_count,
		ChmOnbSCount: chmOnb_count,
		KhnAvaSlide:  khnAva_slide,
		KhnOnbSlide:  khnOnb_slide,
		KhnAvaSCount: khnAva_count,
		KhnOnbSCount: khnOnb_count,
		HdyAvaSlide:  hdyAva_slide,
		HdyOnbSlide:  hdyOnb_slide,
		HdyAvaSCount: hdyAva_count,
		HdyOnbSCount: hdyOnb_count,
	}

	result := models.CountStaffDashBoard{
		Total: staff_total,
		Site:  staff_center,
	}

	return &result, nil
}

func (s staffService) SrvGetStaff() ([]models.StaffDashBoard, error) {
	dashboard_result, err := s.staffRepo.GetStaff()
	if err != nil {
		return nil, err
	}
	return dashboard_result, nil
}
