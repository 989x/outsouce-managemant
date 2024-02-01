package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"osm/api/database"
	"osm/api/helpers"
	"osm/api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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

/*>>>>>>>>>>     Staffs     <<<<<<<<<<*/

var staffSLide []models.Staff

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

	if err := query_result.All(context, &staffSLide); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	return helpers.JsonResponse(c, nil, 200, staffSLide, "Success")
}

func GetFillterStaff(c *fiber.Ctx) error {
	collection := database.MgConn.Db.Collection("staff_jobs")
	context := database.MgConn.Ctx

	date_format_string := "2006-01-02"
	date_query := c.Query("date", time.Now().Add(-24*time.Hour).Format(date_format_string))
	date_nPlus, err := time.Parse(date_format_string, date_query)

	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	var fillterBody models.SearchStaff
	if err := c.BodyParser(&fillterBody); err != nil {
		return helpers.JsonResponse(c, err, 200, nil, "Fail")
	}

	pipeline := helpers.GetAllStaff(date_nPlus)

	query_result, err := collection.Aggregate(context, pipeline)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	if err := query_result.All(context, &staffSLide); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	var searchStaffResults []models.Staff

	searchBar := helpers.SplitParser(fillterBody.Search)
	searchCenter := helpers.SplitParser(fillterBody.Center)
	fmt.Println(fillterBody.Center)
	searchAvailable := helpers.SplitParser(fillterBody.Available)
	searchStatus := helpers.SplitParser(fillterBody.Status)
	searchTeam := helpers.SplitParser(fillterBody.Team)
	searchOutsource := helpers.SplitParser(fillterBody.Outsource)
	searchStatusSite := helpers.SplitParser(fillterBody.Status_site)
	searchSkill := helpers.SplitParser(fillterBody.Skill)

	for index, _ := range staffSLide {
		searchBarResult := false
		searchSkillResult := false

		if searchBar[0] == "" ||
			searchBar[0] == staffSLide[index].ID ||
			searchBar[0] == staffSLide[index].Fname ||
			searchBar[0] == staffSLide[index].Lname ||
			searchBar[0] == staffSLide[index].Nname {
			searchBarResult = true
		}

		for skillIndex, _ := range staffSLide[index].Skill {
			for serachSkillIndex, _ := range searchSkill {
				if staffSLide[index].Skill[skillIndex].Skill == searchSkill[serachSkillIndex] ||
					searchSkill[0] == "All" ||
					searchSkill[0] == "" {
					searchSkillResult = true
					break
				}
			}
		}

		if searchBarResult && searchSkillResult &&
			helpers.SearchCodition(searchCenter, staffSLide[index].Center) &&
			helpers.SearchCodition(searchAvailable, staffSLide[index].Available) &&
			helpers.SearchCodition(searchStatus, staffSLide[index].Status) &&
			helpers.SearchCodition(searchTeam, staffSLide[index].Team) &&
			helpers.SearchCodition(searchOutsource, staffSLide[index].Outsource) &&
			helpers.SearchCodition(searchStatusSite, staffSLide[index].StatusSite) {
			searchStaffResults = append(searchStaffResults, staffSLide[index])
		}
	}

	return helpers.JsonResponse(c, nil, 200, searchStaffResults, "Success")
}

func AddStaff(c *fiber.Ctx) error {
	collectionStaff := database.MgConn.Db.Collection("staffs")
	collectionStaffJob := database.MgConn.Db.Collection("staff_jobs")
	_ = collectionStaffJob
	context := database.MgConn.Ctx

	var staffBody models.NewStaffBody

	if err := c.BodyParser(&staffBody); err != nil {
		return helpers.JsonResponse(c, err, 200, nil, "Fail")
	}

	query := bson.D{{Key: "id", Value: staffBody.ID}}

	queryResult := collectionStaff.FindOne(context, query)

	var exists []models.NewStaffSave
	queryResult.Decode(&exists)

	if len(exists) > 1 {
		err := errors.New("Staff ID has duplicate.")
		return helpers.JsonResponse(c, err, 200, nil, "Fail")
	}

	type StaffRA struct {
		ObjectID string `json:"objectId"`
		Type     string `json:"type"`
		Image    bool   `json:"image"`
		Search   int    `json:"search"`
	}

	checkSrattRa := StaffRA{
		ObjectID: "",
		Type:     "",
		Image:    false,
		Search:   staffBody.ID,
	}

	jsonConv, err := json.Marshal(checkSrattRa)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	buffer := bytes.NewBuffer(jsonConv)

	client := &http.Client{}

	httpRequest, err := http.NewRequest("POST", "https://one.th/api/oauth/getpwd", buffer)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	requestClient, err := client.Do(httpRequest)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	defer requestClient.Body.Close()

	httpResult, err := ioutil.ReadAll(requestClient.Body)

	var staffRaResult models.StaffRaResult

	if err = json.Unmarshal(httpResult, &staffRaResult); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	if staffRaResult.Count < 1 {
		err := errors.New("Not Found AccountID Staff.")
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	staffInsertModel := models.NewStaffSave{
		Email:             staffBody.Email,
		Phone:             staffBody.Phone,
		Active:            staffBody.Active,
		IsTransfer:        staffBody.IsTransfer,
		LastActiveDate:    time.Now(),
		Skill:             staffBody.Skill,
		Certificate:       staffBody.Certificate,
		ID:                staffBody.ID,
		Prefix:            staffBody.Prefix,
		Fname:             staffBody.Fname,
		Lname:             staffBody.Lname,
		Nname:             staffBody.Nname,
		Center:            staffBody.Center,
		Team:              staffBody.Team,
		StartDate:         time.Now(),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		V:                 0,
		AccountID:         staffBody.AccountID,
		OneEmail:          staffBody.OneEmail,
		LeaveDetail:       staffBody.LeaveDetail,
		ResignDescription: staffBody.ResignDescription,
	}

	StaffInsertResult, err := collectionStaff.InsertOne(context, staffInsertModel)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	_ = StaffInsertResult

	query = bson.D{{Key: "_id", Value: StaffInsertResult.InsertedID}}
	queryResult = collectionStaff.FindOne(c.Context(), query)

	var staffLastInsert models.NewStaffSave
	queryResult.Decode(&staffLastInsert)

	staffJobInsertModel := models.NewStaffJob{
		FinishJobsDate: nil,
		AcceptJob:      false,
		UserID:         staffLastInsert.Obj_ID,
		JobID:          nil,
		StartJobsDate:  time.Now(),
		Status:         "Training",
		Available:      "Available",
		Outsource:      "ยังไม่ได้รับงาน",
		Matchjob:       "ว่าง",
		AddressOnsite:  "ตามต้นสังกัด",
		StatusSite:     "Offsite",
		Note:           "สถานะ เริ่มต้น",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		V:              0,
	}

	StaffJobInsertResult, err := collectionStaff.InsertOne(context, staffJobInsertModel)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	query = bson.D{{Key: "_id", Value: StaffJobInsertResult.InsertedID}}
	queryResult = collectionStaff.FindOne(c.Context(), query)

	var staffJobLastInsert models.NewStaffJob
	queryResult.Decode(&staffJobLastInsert)

	staffInsertResult := models.StaffInsertResult{
		NewStaff:    staffLastInsert,
		NewStassJob: staffJobLastInsert,
	}

	return helpers.JsonResponse(c, nil, 200, staffInsertResult, "Success")
}

func GetStaffById(c *fiber.Ctx) error {
	userId := c.Params("id")

	fmt.Println(userId)

	var staffResults []models.Staff

	for index, _ := range staffSLide {
		if staffSLide[index].UserID == userId {
			staffResults = append(staffResults, staffSLide[index])
			break
		}
	}

	var staffResult models.Staff
	staffResult = staffResults[0]

	return helpers.JsonResponse(c, nil, 200, staffResult, "Success")
}

func GetStaffView(c *fiber.Ctx) error {
	collection := database.MgConn.Db.Collection("staffs")
	context := database.MgConn.Ctx

	paramsUser := c.Params("id")
	_id, err := primitive.ObjectIDFromHex(paramsUser)
	if err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}
	var staff models.StaffGetForUpdate

	query := bson.D{{Key: "_id", Value: _id}}
	if err := collection.FindOne(context, query).Decode(&staff); err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}

	return helpers.JsonResponse(c, nil, 200, staff, "Success")
}

func GetStaffJobView(c *fiber.Ctx) error {
	collection := database.MgConn.Db.Collection("staff_jobs")
	context := database.MgConn.Ctx

	paramsUser := c.Params("id")
	userId, err := primitive.ObjectIDFromHex(paramsUser)
	if err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}

	query := bson.D{{Key: "user_id", Value: userId}}
	queryResult, err := collection.Find(context, query)
	if err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}

	var staffJob []models.StaffJobGetForUpdate

	if err := queryResult.All(context, &staffJob); err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}

	return helpers.JsonResponse(c, nil, 200, staffJob, "Success")
}

func UpdateStaff(c *fiber.Ctx) error {
	collection := database.MgConn.Db.Collection("staffs")
	context := database.MgConn.Ctx

	var staff models.StaffGetForUpdate
	if err := c.BodyParser(&staff); err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}

	staff.UpdatedAt = time.Now()

	paramsUser := c.Params("id")
	_id, err := primitive.ObjectIDFromHex(paramsUser)
	if err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}

	query := bson.D{{Key: "_id", Value: _id}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "email", Value: staff.Email},
				{Key: "phone", Value: staff.Phone},
				{Key: "skill", Value: staff.Skill},
				{Key: "prefix", Value: staff.Prefix},
				{Key: "fname", Value: staff.Fname},
				{Key: "lname", Value: staff.Lname},
				{Key: "nname", Value: staff.Nname},
				{Key: "center", Value: staff.Center},
				{Key: "team", Value: staff.Team},
				{Key: "updatedAt", Value: staff.UpdatedAt},
			},
		},
	}

	if err := collection.FindOneAndUpdate(context, query, update).Err(); err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}

	return helpers.JsonResponse(c, nil, 200, nil, "Success")
}
