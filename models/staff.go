package models

import "time"

type Staff struct {
	Email          []interface{} `json:"email"`
	Phone          []interface{} `json:"phone"`
	Active         bool          `json:"active"`
	IsTransfer     bool          `json:"isTransfer"`
	LastActiveDate time.Time     `json:"last_active_date"`
	Skill          []struct {
		Skill string `json:"skill"`
		Level int    `json:"level"`
	} `json:"skill"`
	Certificate       []interface{} `json:"certificate"`
	ID                string        `json:"id"`
	Prefix            string        `json:"prefix"`
	Fname             string        `json:"fname"`
	Lname             string        `json:"lname"`
	Nname             string        `json:"nname"`
	Center            string        `json:"center"`
	Team              string        `json:"team"`
	StartDate         time.Time     `json:"start_date"`
	CreatedAt         time.Time     `json:"createdAt"`
	UpdatedAt         time.Time     `json:"updatedAt"`
	AccountID         string        `json:"account_id"`
	OneEmail          string        `json:"one_email"`
	LeaveDetail       interface{}   `json:"leaveDetail"`
	ResignDescription string        `json:"resign_description"`
}

type StaffResponse struct {
	Obj_ID         string        `json:"_id,omitempty" bson:"_id,omitempty"`
	Email          []interface{} `json:"email" bson:"email,omitempty"`
	Phone          []interface{} `json:"phone" bson:"phone,omitempty"`
	Active         bool          `json:"active" bson:"active,omitempty"`
	IsTransfer     bool          `json:"isTransfer" bson:"isTransfer,omitempty"`
	LastActiveDate time.Time     `json:"last_active_date" bson:"last_active_date,omitempty"`
	Skill          []struct {
		Skill string `json:"skill" bson:"skill,omitempty"`
		Level int    `json:"level" bson:"level,omitempty"`
	} `json:"skill"`
	Certificate       []interface{} `json:"certificate" bson:"certificate,omitempty"`
	ID                string        `json:"id" bson:"id,omitempty"`
	Prefix            string        `json:"prefix" bson:"prefix,omitempty"`
	Fname             string        `json:"fname" bson:"fname,omitempty"`
	Lname             string        `json:"lname" bson:"lname,omitempty"`
	Nname             string        `json:"nname" bson:"nname,omitempty"`
	Center            string        `json:"center" bson:"center,omitempty"`
	Team              string        `json:"team" bson:"team,omitempty"`
	AccountID         string        `json:"account_id" bson:"account_id,omitempty"`
	OneEmail          string        `json:"one_email" bson:"one_email,omitempty"`
	LeaveDetail       interface{}   `json:"leaveDetail" bson:"leaveDetail,omitempty"`
	ResignDescription string        `json:"resign_description" bson:"resign_description,omitempty"`
}

type StaffJobResponse struct {
	Obj_ID         string      `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID         string      `json:"user_id,omitempty" bson:"user_id,omitempty"`
	StartJobsDate  time.Time   `json:"start_jobs_date"  bson:"start_jobs_date,omitempty"`
	FinishJobsDate interface{} `json:"finish_jobs_date"`
	Status         string      `json:"status" bson:"status,omitempty"`
	Available      string      `json:"available" bson:"available,omitempty"`
	Outsource      string      `json:"outsource" bson:"outsource,omitempty"`
	Matchjob       string      `json:"matchjob" bson:"matchjob,omitempty"`
	AddressOnsite  string      `json:"address_onsite" bson:"address_onsite,omitempty"`
	StatusSite     string      `json:"status_site" bson:"status_site,omitempty"`
	Note           string      `json:"note" bson:"note,omitempty"`
	JobID          interface{} `json:"job_id" bson:"job_id,omitempty"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAt"`
	ID             string      `json:"id" bson:"id,omitempty"`
	Prefix         string      `json:"prefix" bson:"prefix,omitempty"`
	Fname          string      `json:"fname" bson:"fname,omitempty"`
	Lname          string      `json:"lname" bson:"lname,omitempty"`
	Nname          string      `json:"nname" bson:"nname,omitempty"`
	StartDate      time.Time   `json:"start_date" bson:"start_date,omitempty"`
	Skill          []struct {
		Skill string `json:"skill" bson:"skill,omitempty"`
		Level int    `json:"level" bson:"level,omitempty"`
	} `json:"skill"`
	Email          []interface{} `json:"email"`
	Phone          []string      `json:"phone"`
	LeaveDetail    interface{}   `json:"leaveDetail"`
	Active         bool          `json:"active"`
	IsTransfer     bool          `json:"isTransfer"`
	LastActiveDate interface{}   `json:"last_active_date"`
	Center         string        `json:"center"`
	Team           string        `json:"team"`
	Certificate    []interface{} `json:"certificate"`
	AccountID      string        `json:"account_id"`
	JobAll         []struct {
		ID struct {
			Oid string `json:"$oid"`
		} `json:"_id"`
		FinishJobsDate interface{} `json:"finish_jobs_date"`
		AcceptSignJob  bool        `json:"accept_sign_job"`
		AcceptStrtJob  bool        `json:"accept_strt_job"`
		AcceptFinsJob  bool        `json:"accept_fins_job"`
		UserID         struct {
			Oid string `json:"$oid"`
		} `json:"user_id"`
		JobID         interface{} `json:"job_id"`
		StartJobsDate time.Time
		Status        string    `json:"status" bson:"status,omitempty"`
		Available     string    `json:"available" bson:"available,omitempty"`
		Outsource     string    `json:"outsource" bson:"outsource,omitempty"`
		Matchjob      string    `json:"matchjob" bson:"matchjob,omitempty"`
		AddressOnsite string    `json:"address_onsite" bson:"address_onsite,omitempty"`
		StatusSite    string    `json:"status_site" bson:"status_site,omitempty"`
		Note          string    `json:"note" bson:"note,omitempty"`
		CreatedAt     time.Time `json:"createdAt"`
		UpdatedAt     time.Time `json:"updatedAt"`
	} `json:"job_all"`
	CvURL  []interface{} `json:"cv_url"`
	CvName []interface{} `json:"cv_name"`
}

type StaffExistsRequest struct {
	ObjectId string
	Type     string
	Tmage    bool
	search   int64
}

type StaffExists struct {
	Success bool `json:"success"`
	Count   int  `json:"count"`
	Result  []struct {
		ID                  string    `json:"_id"`
		UserID              string    `json:"userId"`
		AccountID           string    `json:"accountId"`
		TitleTh             string    `json:"titleTh"`
		FirstNameTh         string    `json:"firstNameTh"`
		LastNameTh          string    `json:"lastNameTh"`
		NameTh              string    `json:"nameTh"`
		TitleEn             string    `json:"titleEn"`
		FirstNameEn         string    `json:"firstNameEn"`
		LastNameEn          string    `json:"lastNameEn"`
		NameEn              string    `json:"nameEn"`
		Email               string    `json:"email"`
		EmailOneID          string    `json:"emailOneId"`
		NickName            string    `json:"nickName"`
		Tel                 string    `json:"tel"`
		EmployeeID          string    `json:"employeeId"`
		PositionID          string    `json:"positionId"`
		PositionName        string    `json:"positionName"`
		PositionLevel       string    `json:"positionLevel"`
		StartWorkDate       time.Time `json:"startWorkDate"`
		TaxID               string    `json:"taxId"`
		CompanyID           string    `json:"companyId"`
		CompanyFullNameTh   string    `json:"companyFullNameTh"`
		CompanyFullNameEng  string    `json:"companyFullNameEng"`
		CompanyShortNameTh  string    `json:"companyShortNameTh"`
		CompanyShortNameEng string    `json:"companyShortNameEng"`
		Station             string    `json:"station"`
		ContractType        string    `json:"contractType"`
		Resign              bool      `json:"resign"`
		List                []struct {
			ID           string `json:"_id"`
			OrgChartType string `json:"orgChartType"`
			OrgChartName string `json:"orgChartName"`
			CompanyID    string `json:"companyId"`
			CompanyName  string `json:"companyName"`
			Name         string `json:"name"`
		} `json:"list"`
		History       []interface{} `json:"history"`
		LeaderHistory []interface{} `json:"leaderHistory"`
	} `json:"result"`
}

type StaffJobsFind struct {
	ID             string      `json:"_id,omitempty" bson:"_id,omitempty"`
	FinishJobsDate interface{} `json:"finish_jobs_date"`
	AcceptJob      bool        `json:"accept_job"`
	UserID         string      `json:"user_id" bson:"user_id,omitempty"`
	JobID          interface{} `json:"job_id"`
	StartJobsDate  time.Time   `json:"start_jobs_date"  bson:"start_jobs_date,omitempty"`
	Status         string      `json:"status"`
	Available      string      `json:"available"`
	Outsource      string      `json:"outsource"`
	Matchjob       string      `json:"matchjob"`
	AddressOnsite  string      `json:"address_onsite"`
	StatusSite     string      `json:"status_site"`
	Note           string      `json:"note"`
}
