package controllers

import (
	"osm/api/database"
	"osm/api/helpers"
	"osm/api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetStaffDashBoard(c *fiber.Ctx) error {
	collection := database.MgConn.Db.Collection("staff_jobs")
	context := database.MgConn.Ctx

	date_format_string := "2006-01-02"
	date_query := c.Query("date", time.Now().Add(-24*time.Hour).Format(date_format_string))
	date_nPlus, err := time.Parse(date_format_string, date_query)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	var result_total []models.StaffDashBoard

	// --> Prepare Filters
	availableValues := []interface{}{"Available"}
	teamValues := []interface{}{"Dev", "Tester", "UXUI", "Data Sci", "DevOps"}
	statusValues := []interface{}{"ลา", "ลาไม่รับเงินเดือน", "ลาคลอด", "ลาอุปสมบท", "ลารับราชการทหาร", "ลาฌาปนกิจ"}
	centerValues := []interface{}{"ขอนแก่น", "เชียงใหม่", "หาดใหญ่", "กรุงเทพ"}
	filter_devAvanonit := helpers.StaffPrimetiveFilter(availableValues, teamValues, statusValues, centerValues)
	// ____________________________________________Query - Dev Available____________________________________________//
	getDevAvanonit := helpers.StaffPipeLineTotal(filter_devAvanonit, date_nPlus)

	query_result, err := collection.Aggregate(context, getDevAvanonit)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	if err := query_result.All(context, &result_total); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	getDevAvanonit_count := len(result_total) // Total Dev Available
	// ____________________________________________Query - Dev Available____________________________________________//

	// --> Prepare Filters
	availableValues = []interface{}{"On Board"}
	filter_devOnnonitt := helpers.StaffPrimetiveFilter(availableValues, teamValues, statusValues, centerValues)
	//  ____________________________________________Query - Dev On board ____________________________________________//

	getDevOnnonit := helpers.StaffPipeLineTotal(filter_devOnnonitt, date_nPlus)
	query_result, err = collection.Aggregate(context, getDevOnnonit)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	if err := query_result.All(context, &result_total); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	getDevOnnonit_count := len(result_total) // Total Dev On Board
	//  ____________________________________________Query - Dev On board ____________________________________________//

	// --> Prepare Filters
	availableValues = []interface{}{"Available"}
	teamValues = []interface{}{"IT Infra"}
	filter_itAva := helpers.StaffPrimetiveFilter(availableValues, teamValues, statusValues, centerValues)
	//  ____________________________________________Query - It Available ____________________________________________//
	getItAva := helpers.StaffPipeLineTotal(filter_itAva, date_nPlus)
	query_result, err = collection.Aggregate(context, getItAva)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	if err := query_result.All(context, &result_total); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	getItAva_count := len(result_total) // Total It On Available
	//  ____________________________________________Query - It Available ____________________________________________//

	// --> Prepare Filters
	availableValues = []interface{}{"On Board"}
	filter_itOn := helpers.StaffPrimetiveFilter(availableValues, teamValues, statusValues, centerValues)
	// ____________________________________________Query - It On Board____________________________________________//
	getItOn := helpers.StaffPipeLineTotal(filter_itOn, date_nPlus)
	query_result, err = collection.Aggregate(context, getItOn)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	if err := query_result.All(context, &result_total); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	getItOn_count := len(result_total) // Total It On Board
	// ____________________________________________Query - It On Board____________________________________________//

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

	// >> Count BnkAva
	count_target := helpers.GetCountCenterStaff("กรุงเทพ", "Available", date_nPlus)
	query_result, err = collection.Aggregate(context, count_target)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	var bnkAva_slide []models.StaffCenterStatus
	if err := query_result.All(context, &bnkAva_slide); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	bnkAva_count := len(bnkAva_slide) //<< BnkAva Result

	// >> Count BnkOnb
	count_target = helpers.GetCountCenterStaff("กรุงเทพ", "On Board", date_nPlus)
	query_result, err = collection.Aggregate(context, count_target)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	var bnkOnb_slide []models.StaffCenterStatus
	if err := query_result.All(context, &bnkOnb_slide); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	bnkOnb_count := len(bnkOnb_slide) //<< BnkOnb Result

	// >> Count ChmAva
	count_target = helpers.GetCountCenterStaff("เชียงใหม่", "Available", date_nPlus)
	query_result, err = collection.Aggregate(context, count_target)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	var chmAva_slide []models.StaffCenterStatus
	if err := query_result.All(context, &chmAva_slide); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	chmAva_count := len(chmAva_slide) //<< ChmAva Result

	// >> Count ChmOnb
	count_target = helpers.GetCountCenterStaff("เชียงใหม่", "On Board", date_nPlus)
	query_result, err = collection.Aggregate(context, count_target)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	var chmOnb_slide []models.StaffCenterStatus
	if err := query_result.All(context, &chmOnb_slide); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	chmOnb_count := len(chmOnb_slide) //<< ChmOnb Result

	// >> Count KhnAva
	count_target = helpers.GetCountCenterStaff("ขอนแก่น", "Available", date_nPlus)
	query_result, err = collection.Aggregate(context, count_target)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	var khnAva_slide []models.StaffCenterStatus
	if err := query_result.All(context, &khnAva_slide); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	khnAva_count := len(khnAva_slide) //<< KhnAva Result

	// >> Count KhnOnb
	count_target = helpers.GetCountCenterStaff("ขอนแก่น", "On Board", date_nPlus)
	query_result, err = collection.Aggregate(context, count_target)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	var khnOnb_slide []models.StaffCenterStatus
	if err := query_result.All(context, &khnOnb_slide); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	khnOnb_count := len(khnOnb_slide) //<< KhnOnb Result

	// >> Count HdyAva
	count_target = helpers.GetCountCenterStaff("หาดใหญ่", "Available", date_nPlus)
	query_result, err = collection.Aggregate(context, count_target)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	var hdyAva_slide []models.StaffCenterStatus
	if err := query_result.All(context, &hdyAva_slide); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	hdyAva_count := len(hdyAva_slide) //<< HdyAva Result

	// >> Count HdyOnb
	count_target = helpers.GetCountCenterStaff("หาดใหญ่", "On Board", date_nPlus)
	query_result, err = collection.Aggregate(context, count_target)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	var hdyOnb_slide []models.StaffCenterStatus
	if err := query_result.All(context, &hdyOnb_slide); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	hdyOnb_count := len(hdyOnb_slide) //<< HdyAva HdyOnb

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

	project_query := helpers.GetProjectTotal(date_nPlus)
	query_result, err = collection.Aggregate(context, project_query)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	var project_slide []models.StaffGetProject
	if err := query_result.All(context, &project_slide); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	response_data := models.CountStaffDashBoard{
		Total:   staff_total,
		Site:    staff_center,
		Project: project_slide,
	}

	return helpers.JsonResponse(c, nil, 200, response_data, "Success")
}

func GetStaffByJobName(c *fiber.Ctx) error {
	collection := database.MgConn.Db.Collection("staff_jobs")
	context := database.MgConn.Ctx

	projectId := c.Params("project")
	objID, err := primitive.ObjectIDFromHex(projectId)

	date_format_string := "2006-01-02"
	date_query := c.Query("date", time.Now().Add(-24*time.Hour).Format(date_format_string))
	date_nPlus, err := time.Parse(date_format_string, date_query)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	participant_pipeline := helpers.GetStaffParticipant(date_nPlus, objID)

	query_result, err := collection.Aggregate(context, participant_pipeline)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	var participant_slide []models.StaffParticipant

	if err := query_result.All(context, &participant_slide); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	return helpers.JsonResponse(c, nil, 200, participant_slide, "Success")
}

func GetStaff(c *fiber.Ctx) error {
	collection := database.MgConn.Db.Collection("staff_jobs")
	context := database.MgConn.Ctx

	date_format_string := "2006-01-02"
	date_query := c.Query("date", time.Now().Add(-24*time.Hour).Format(date_format_string))
	date_nPlus, err := time.Parse(date_format_string, date_query)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	pipeline := helpers.GetAllStaff(date_nPlus)

	query_result, err := collection.Aggregate(context, pipeline)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	var staff_slide []models.Staff

	if err := query_result.All(context, &staff_slide); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	return helpers.JsonResponse(c, nil, 200, staff_slide, "Success")
}
