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

type Get_OneAccount struct {
	ID                  string      `json:"id"`
	FirstNameTh         string      `json:"first_name_th"`
	MiddleNameTh        interface{} `json:"middle_name_th"`
	LastNameTh          string      `json:"last_name_th"`
	FirstNameEng        string      `json:"first_name_eng"`
	MiddleNameEng       interface{} `json:"middle_name_eng"`
	LastNameEng         string      `json:"last_name_eng"`
	SpecialTitleNameTh  string      `json:"special_title_name_th"`
	AccountTitleTh      string      `json:"account_title_th"`
	SpecialTitleNameEng interface{} `json:"special_title_name_eng"`
	AccountTitleEng     string      `json:"account_title_eng"`
	IDCardType          string      `json:"id_card_type"`
	IDCardNum           string      `json:"id_card_num"`
	HashIDCardNum       string      `json:"hash_id_card_num"`
	AccountCategory     string      `json:"account_category"`
	AccountSubCategory  string      `json:"account_sub_category"`
	ThaiEmail           string      `json:"thai_email"`
	ThaiEmail2          string      `json:"thai_email2"`
	ThaiEmail3          interface{} `json:"thai_email3"`
	StatusCd            string      `json:"status_cd"`
	BirthDate           string      `json:"birth_date"`
	StatusDt            string      `json:"status_dt"`
	RegisterDt          string      `json:"register_dt"`
	AddressID           interface{} `json:"address_id"`
	CreatedAt           string      `json:"created_at"`
	CreatedBy           string      `json:"created_by"`
	UpdatedAt           string      `json:"updated_at"`
	UpdatedBy           string      `json:"updated_by"`
	Reason              interface{} `json:"reason"`
	TelNo               interface{} `json:"tel_no"`
	NameOnDocumentTh    interface{} `json:"name_on_document_th"`
	NameOnDocumentEng   interface{} `json:"name_on_document_eng"`
	BlockchainFlg       interface{} `json:"blockchain_flg"`
	NicknameEng         string      `json:"nickname_eng"`
	NicknameTh          string      `json:"nickname_th"`
	TrustLevel          interface{} `json:"trust_level"`
	Mobile              []struct {
		ID        string      `json:"id"`
		MobileNo  string      `json:"mobile_no"`
		CreatedAt string      `json:"created_at"`
		CreatedBy string      `json:"created_by"`
		UpdatedAt string      `json:"updated_at"`
		UpdatedBy string      `json:"updated_by"`
		DeletedAt interface{} `json:"deleted_at"`
		Pivot     struct {
			AccountID             string      `json:"account_id"`
			MobileID              string      `json:"mobile_id"`
			CreatedAt             string      `json:"created_at"`
			UpdatedAt             string      `json:"updated_at"`
			StatusCd              string      `json:"status_cd"`
			PrimaryFlg            string      `json:"primary_flg"`
			MobileConfirmFlg      string      `json:"mobile_confirm_flg"`
			MobileConfirmDt       string      `json:"mobile_confirm_dt"`
			MobileLoginConfirmFlg string      `json:"mobile_login_confirm_flg"`
			MobileLoginConfirmDt  interface{} `json:"mobile_login_confirm_dt"`
			Type                  string      `json:"type"`
		} `json:"pivot"`
	} `json:"mobile"`
	Email []struct {
		ID        string      `json:"id"`
		Email     string      `json:"email"`
		CreatedAt string      `json:"created_at"`
		CreatedBy string      `json:"created_by"`
		UpdatedAt string      `json:"updated_at"`
		UpdatedBy string      `json:"updated_by"`
		DeletedAt interface{} `json:"deleted_at"`
		Pivot     struct {
			AccountID            string      `json:"account_id"`
			EmailID              string      `json:"email_id"`
			CreatedAt            string      `json:"created_at"`
			UpdatedAt            string      `json:"updated_at"`
			StatusCd             string      `json:"status_cd"`
			PrimaryFlg           string      `json:"primary_flg"`
			EmailConfirmFlg      string      `json:"email_confirm_flg"`
			EmailConfirmDt       string      `json:"email_confirm_dt"`
			EmailLoginConfirmFlg string      `json:"email_login_confirm_flg"`
			EmailLoginConfirmDt  interface{} `json:"email_login_confirm_dt"`
		} `json:"pivot"`
	} `json:"email"`
	Address          []interface{} `json:"address"`
	AccountAttribute []interface{} `json:"account_attribute"`
	Status           string        `json:"status"`
	LastUpdate       string        `json:"last_update"`
}
