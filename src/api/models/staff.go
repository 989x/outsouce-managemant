package models

import "time"

type CountStaffDashBoard struct {
	Total   interface{}
	Site    interface{}
	Project interface{}
}

type StaffTotal struct {
	All          int
	AllAvailable int
	AllOnBoard   int
	DevOnBoard   int
	DevAvailable int
	AllDev       int
	ItOnBoard    int
	ItAvailable  int
	AllIt        int
}

type StaffCountCenter struct {
	BnkAvaSlide  interface{}
	BnkOnbSlide  interface{}
	BnkAvaSCount int
	BnkOnbSCount int
	ChmAvaSlide  interface{}
	ChmOnbSlide  interface{}
	ChmAvaSCount int
	ChmOnbSCount int
	KhnAvaSlide  interface{}
	KhnOnbSlide  interface{}
	KhnAvaSCount int
	KhnOnbSCount int
	HdyAvaSlide  interface{}
	HdyOnbSlide  interface{}
	HdyAvaSCount int
	HdyOnbSCount int
}

type StaffGetProject struct {
	ID                 string `json:"_id" bson:"_id,omitempty"`
	ProjectName        string `json:"x" bson:"projectName,omitempty"`
	ProjectParticipant int    `json:"y" bson:"projectParticipant,omitempty"`
}

type StaffProjectApex struct {
	X string `json:"x"`
	Y int    `json:"y"`
}

type StaffDashBoard struct {
	Obj_ID         string      `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID         string      `json:"user_id" bson:"user_id,omitempty"`
	StartJobsDate  time.Time   `json:"start_jobs_date" bson:"start_jobs_date,omitempty"`
	FinishJobsDate time.Time   `json:"finish_jobs_date" bson:"finish_jobs_date,omitempty"`
	Status         string      `json:"status" bson:"status,omitempty"`
	Available      string      `json:"available" bson:"available,omitempty"`
	Outsource      string      `json:"outsource" bson:"outsource,omitempty"`
	Matchjob       string      `json:"matchjob" bson:"matchjob,omitempty"`
	AddressOnsite  string      `json:"address_onsite" bson:"address_onsite,omitempty"`
	StatusSite     string      `json:"status_site" bson:"status_site,omitempty"`
	Note           string      `json:"note" bson:"note,omitempty"`
	CreatedAt      time.Time   `json:"createdAt"  bson:"createdAt,omitempty"`
	UpdatedAt      time.Time   `json:"updatedAt"  bson:"updatedAt,omitempty"`
	ID             string      `json:"id" bson:"id,omitempty"`
	Fname          string      `json:"fname" bson:"fname,omitempty"`
	Lname          string      `json:"lname" bson:"lname,omitempty"`
	Nname          string      `json:"nname" bson:"nname,omitempty"`
	StartDate      time.Time   `json:"start_date" bson:"start_date,omitempty"`
	Active         bool        `json:"active" bson:"active,omitempty"`
	IsTransfer     bool        `json:"isTransfer" bson:"isTransfer,omitempty"`
	LastActiveDate interface{} `json:"last_active_date" bson:"last_active_date,omitempty"`
	Center         string      `json:"center" bson:"center,omitempty"`
	Team           string      `json:"team" bson:"team,omitempty"`
	AccountID      string      `json:"account_id" bson:"account_id,omitempty"`
}

type StaffCenterStatus struct {
	Obj_ID         string      `json:"_id" bson:"_id,omitempty"`
	UserID         string      `json:"user_id" bson:"user_id,omitempty"`
	StartJobsDate  time.Time   `json:"start_jobs_date" bson:"start_jobs_date,omitempty"`
	FinishJobsDate time.Time   `json:"finish_jobs_date" bson:"finish_jobs_date,omitempty"`
	Matchjob       string      `json:"matchjob" bson:"matchjob,omitempty"`
	Status         string      `json:"status" bson:"status,omitempty"`
	AddressOnsite  string      `json:"address_onsite" bson:"address_onsite,omitempty"`
	StatusSite     string      `json:"status_site" bson:"status_site,omitempty"`
	CreatedAt      time.Time   `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt      time.Time   `json:"updatedAt" bson:"updatedAt,omitempty"`
	Available      string      `json:"available" bson:"available,omitempty"`
	Outsource      string      `json:"outsource" bson:"outsource,omitempty"`
	Note           string      `json:"note" bson:"note,omitempty"`
	JobID          string      `json:"job_id" bson:"job_id,omitempty"`
	ID             string      `json:"id" bson:"id,omitempty"`
	Fname          string      `json:"fname" bson:"fname,omitempty"`
	Lname          string      `json:"lname" bson:"lname,omitempty"`
	Nname          string      `json:"nname" bson:"nname,omitempty"`
	StartDate      time.Time   `json:"start_date" bson:"start_date,omitempty"`
	Active         bool        `json:"active" bson:"active,omitempty"`
	IsTransfer     bool        `json:"isTransfer" bson:"isTransfer,omitempty"`
	LastActiveDate interface{} `json:"last_active_date" bson:"last_active_date,omitempty"`
	Center         string      `json:"center" bson:"center,omitempty"`
	Team           string      `json:"team" bson:"team,omitempty"`
	AccountID      string      `json:"account_id" bson:"account_id,omitempty"`
}

type StaffParticipant struct {
	Obj_ID         string      `json:"_id" bson:"_id,omitempty"`
	UserID         string      `json:"user_id" bson:"user_id,omitempty"`
	StartJobsDate  time.Time   `json:"start_jobs_date" bson:"start_jobs_date,omitempty"`
	FinishJobsDate time.Time   `json:"finish_jobs_date" bson:"finish_jobs_date,omitempty"`
	Status         string      `json:"status" bson:"status,omitempty"`
	Matchjob       string      `json:"matchjob" bson:"matchjob,omitempty"`
	AddressOnsite  string      `json:"address_onsite" bson:"address_onsite,omitempty"`
	StatusSite     string      `json:"status_site" bson:"status_site,omitempty"`
	CreatedAt      time.Time   `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt      time.Time   `json:"updatedAt" bson:"updatedAt,omitempty"`
	Available      string      `json:"available" bson:"available,omitempty"`
	Outsource      string      `json:"outsource" bson:"outsource,omitempty"`
	Note           string      `json:"note" bson:"note,omitempty"`
	JobID          string      `json:"job_id" bson:"job_id,omitempty"`
	ID             string      `json:"id" bson:"id,omitempty"`
	Fname          string      `json:"fname" bson:"fname,omitempty"`
	Lname          string      `json:"lname" bson:"lname,omitempty"`
	Nname          string      `json:"nname" bson:"nname,omitempty"`
	StartDate      time.Time   `json:"start_date" bson:"start_date,omitempty"`
	Active         bool        `json:"active" bson:"active,omitempty"`
	IsTransfer     bool        `json:"isTransfer" bson:"isTransfer,omitempty"`
	LastActiveDate interface{} `json:"last_active_date" bson:"last_active_date,omitempty"`
	Center         string      `json:"center" bson:"center,omitempty"`
	Team           string      `json:"team" bson:"team,omitempty"`
	AccountID      string      `json:"account_id" bson:"account_id,omitempty"`
}

type NewStaffBody struct {
	Email      []interface{} `json:"email"`
	Phone      []interface{} `json:"phone"`
	Active     bool          `json:"active"`
	IsTransfer bool          `json:"isTransfer"`
	Skill      []struct {
		Skill string `json:"skill"`
		Level int    `json:"level"`
	} `json:"skill"`
	Certificate       []interface{} `json:"certificate"`
	ID                int           `json:"id"`
	Prefix            string        `json:"prefix"`
	Fname             string        `json:"fname"`
	Lname             string        `json:"lname"`
	Nname             string        `json:"nname"`
	Center            string        `json:"center"`
	Team              string        `json:"team"`
	AccountID         string        `json:"account_id"`
	OneEmail          string        `json:"one_email"`
	LeaveDetail       interface{}   `json:"leaveDetail"`
	ResignDescription string        `json:"resign_description"`
}

type NewStaffSave struct {
	Obj_ID         string        `json:"_id" bson:"_id,omitempty"`
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
	ID                int           `json:"id"`
	Prefix            string        `json:"prefix"`
	Fname             string        `json:"fname"`
	Lname             string        `json:"lname"`
	Nname             string        `json:"nname"`
	Center            string        `json:"center"`
	Team              string        `json:"team"`
	StartDate         time.Time     `json:"start_date"`
	CreatedAt         time.Time     `json:"createdAt"`
	UpdatedAt         time.Time     `json:"updatedAt"`
	V                 int           `json:"__v"`
	AccountID         string        `json:"account_id"`
	OneEmail          string        `json:"one_email"`
	LeaveDetail       interface{}   `json:"leaveDetail"`
	ResignDescription string        `json:"resign_description"`
}

type Staff struct {
	Obj_ID         string      `json:"_id" bson:"_id,omitempty"`
	UserID         string      `json:"user_id" bson:"user_id,omitempty"`
	StartJobsDate  time.Time   `json:"start_jobs_date" bson:"start_jobs_date,omitempty"`
	FinishJobsDate time.Time   `json:"finish_jobs_date" bson:"finish_jobs_date,omitempty"`
	Status         string      `json:"status" bson:"status,omitempty"`
	Matchjob       string      `json:"matchjob" bson:"matchjob,omitempty"`
	AddressOnsite  string      `json:"address_onsite" bson:"address_onsite,omitempty"`
	StatusSite     string      `json:"status_site" bson:"status_site,omitempty"`
	CreatedAt      time.Time   `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt      time.Time   `json:"updatedAt" bson:"updatedAt,omitempty"`
	Available      string      `json:"available" bson:"available,omitempty"`
	Outsource      string      `json:"outsource" bson:"outsource,omitempty"`
	Note           string      `json:"note" bson:"note,omitempty"`
	JobID          string      `json:"job_id" bson:"job_id,omitempty"`
	ID             string      `json:"id" bson:"id,omitempty"`
	Fname          string      `json:"fname" bson:"fname,omitempty"`
	Lname          string      `json:"lname" bson:"lname,omitempty"`
	Nname          string      `json:"nname" bson:"nname,omitempty"`
	StartDate      time.Time   `json:"start_date" bson:"start_date,omitempty"`
	Active         bool        `json:"active" bson:"active,omitempty"`
	IsTransfer     bool        `json:"isTransfer" bson:"isTransfer,omitempty"`
	LastActiveDate interface{} `json:"last_active_date" bson:"last_active_date,omitempty"`
	Center         string      `json:"center" bson:"center,omitempty"`
	Team           string      `json:"team" bson:"team,omitempty"`
	AccountID      string      `json:"account_id" bson:"account_id,omitempty"`
	Skill          []struct {
		Skill string `json:"skill" bson:"skill,omitempty"`
		Level int    `json:"level" bson:"level,omitempty"`
	} `json:"skill" bson:"skill,omitempty"`
}

type SearchStaff struct {
	Search      string `json:"search"`
	Center      string `json:"center"`
	Available   string `json:"available"`
	Status      string `json:"status"`
	Team        string `json:"team"`
	Outsource   string `json:"outsource"`
	Status_site string `json:"status_site"`
	Skill       string `json:"skill"`
}

type StaffRaResult struct {
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

type NewStaffJob struct {
	FinishJobsDate interface{} `json:"finish_jobs_date"`
	AcceptJob      bool        `json:"accept_job"`
	UserID         string      `json:"user_id"`
	JobID          interface{} `json:"job_id"`
	StartJobsDate  time.Time   `json:"start_jobs_date"`
	Status         string      `json:"status"`
	Available      string      `json:"available"`
	Outsource      string      `json:"outsource"`
	Matchjob       string      `json:"matchjob"`
	AddressOnsite  string      `json:"address_onsite"`
	StatusSite     string      `json:"status_site"`
	Note           string      `json:"note"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAt"`
	V              int         `json:"__v"`
}

type StaffInsertResult struct {
	NewStaff    interface{}
	NewStassJob interface{}
}

type StaffGetForUpdate struct {
	UserID string        `json:"user_id" bson:"_id,omitempty"`
	Email  []interface{} `json:"email" bson:"email,omitempty"`
	Phone  []interface{} `json:"phone" bson:"phone,omitempty"`
	Skill  []struct {
		Skill string `json:"skill" bson:"skill,omitempty"`
		Level int    `json:"level" bson:"level,omitempty"`
	} `json:"skill" bson:"skill,omitempty"`
	ID        string    `json:"id" bson:"id,omitempty"`
	Prefix    string    `json:"prefix" bson:"prefix,omitempty"`
	Fname     string    `json:"fname" bson:"fname,omitempty"`
	Lname     string    `json:"lname" bson:"lname,omitempty"`
	Nname     string    `json:"nname" bson:"nname,omitempty"`
	Center    string    `json:"center" bson:"center,omitempty"`
	Team      string    `json:"team" bson:"team,omitempty"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt,omitempty"`
}
type StaffSkill struct {
	Skill string `json:"skill" bson:"skill,omitempty"`
	Level int    `json:"level" bson:"level,omitempty"`
}

type StaffJobGetForUpdate struct {
	FinishJobsDate interface{} `json:"finish_jobs_date" bson:"finish_jobs_date,omitempty"`
	UserID         string      `json:"user_id" bson:"user_id,omitempty"`
	JobID          interface{} `json:"job_id" bson:"job_id,omitempty"`
	StartJobsDate  time.Time   `json:"start_jobs_date" bson:"start_jobs_date,omitempty"`
	Status         string      `json:"status" bson:"status,omitempty"`
	Available      string      `json:"available" bson:"available,omitempty"`
	Outsource      string      `json:"outsource" bson:"outsource,omitempty"`
	Matchjob       string      `json:"matchjob" bson:"matchjob,omitempty"`
	AddressOnsite  string      `json:"address_onsite" bson:"address_onsite,omitempty"`
	StatusSite     string      `json:"status_site" bson:"status_site,omitempty"`
}
