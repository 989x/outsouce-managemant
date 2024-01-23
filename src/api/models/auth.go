package models

// Redirect struct
type Redirect_response struct {
	Key_1 string `json:"Key_1"`
	Key_2 string `json:"Key_2"`
}

type Login_Find_result struct {
	Role            string        `json:"role" bson:"role,omitempty"`
	Email           []string      `json:"email" bson:"email,omitempty"`
	Phone           []interface{} `json:"phone" bson:"phone,omitempty"`
	ApproveStatus   bool          `json:"approve_status" bson:"approve_status,omitempty"`
	StaffID         interface{}   `json:"staff_id" bson:"staff_id,omitempty"`
	AccountID       string        `json:"account_id" bson:"account_id,omitempty"`
	FirstNameTh     string        `json:"first_name_th" bson:"first_name_th,omitempty"`
	LastNameTh      string        `json:"last_name_th" bson:"last_name_th,omitempty"`
	FirstNameEng    string        `json:"first_name_eng" bson:"first_name_eng,omitempty"`
	LastNameEng     string        `json:"last_name_eng" bson:"last_name_eng,omitempty"`
	AccountTitleTh  string        `json:"account_title_th" bson:"account_title_th,omitempty"`
	AccountTitleEng string        `json:"account_title_eng" bson:"account_title_eng,omitempty"`
}

type Admin_login_response struct {
	Usernsme    string
	AccountId   string
	AccessToken string
}

// Login struct
type Login_body struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Login_request struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

type Login_response struct {
	TokenType      string `json:"token_type"`
	ExpiresIn      int    `json:"expires_in"`
	AccessToken    string `json:"access_token"`
	RefreshToken   string `json:"refresh_token"`
	ExpirationDate string `json:"expiration_date"`
	AccountID      string `json:"account_id"`
	Result         string `json:"result"`
	Username       string `json:"username"`
	LoginBy        string `json:"login_by"`
}
