package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"osm/api/models"
	repository "osm/api/repository/auth_repo"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	authRepo repository.AuthRepository
}

func NewAuthRepository(authRepo repository.AuthRepository) AuthService {
	return authService{authRepo: authRepo}
}

func (s authService) SrvLogin(login_body *models.Login_body) (*models.Login_response, error) {
	godotenv.Load(".env")

	validate := validator.New()
	err_validate := validate.Struct(login_body)
	if err_validate != nil {
		return nil, err_validate
	}

	login_request := models.Login_request{
		GrantType:    "password",
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Username:     login_body.Username,
		Password:     login_body.Password,
	}

	json_conv, err := json.Marshal(login_request)
	if err != nil {
		return nil, errors.New("form login mid match")
	}

	buffer := bytes.NewBuffer(json_conv)

	request_call, err := s.authRepo.RepGetPwd(buffer)
	if err != nil {
		return nil, errors.New("Get pwd invalid")
	}

	request_call.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	request_client, err := client.Do(request_call)
	if err != nil {
		return nil, errors.New("not compile request client")
	}

	fmt.Println(request_client)

	defer request_client.Body.Close()

	result, err := ioutil.ReadAll(request_client.Body)

	var login_reqponse models.Login_response
	err = json.Unmarshal(result, &login_reqponse)
	if err != nil {
		return nil, errors.New("can't convert json format")
	}

	return &login_reqponse, nil
}

func (s authService) SrvAdminLogin(login_body *models.Login_body) (*models.Admin_login_response, error) {
	validate := validator.New()
	err_validate := validate.Struct(login_body)
	if err_validate != nil {
		return nil, err_validate
	}

	if login_body.Username == "outsource_admin" {

		result_login, err := s.authRepo.RepFindAdmin()
		if err != nil {
			return nil, err
		}
		result_login.AccountTitleEng = "Mr"
		result_login.AccountTitleTh = "นาย"

		hashpass := "$2b$13$u0PZb2mrgGeeByw3L5dJCuV/u3JP2fRHSCyAAjPI8vi2dtQw5XegS"
		if err = bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(login_body.Password)); err != nil {
			return nil, err
		}

		act_token, err := CreateToken(result_login, "JWT_SECRET", 24)
		response_login := models.Admin_login_response{
			Usernsme:    "admin",
			AccountId:   result_login.AccountID,
			AccessToken: act_token,
		}

		return &response_login, nil
	} else {
		return nil, nil
	}
}

func CreateToken(staffResult *models.Login_Find_result, env string, exp int) (string, error) {
	cliams := jwt.MapClaims{
		"iss":               staffResult.AccountID,
		"id":                staffResult.StaffID,
		"fname_eng":         staffResult.FirstNameEng,
		"lname_eng":         staffResult.LastNameEng,
		"fname_th":          staffResult.FirstNameTh,
		"lname_th":          staffResult.LastNameTh,
		"account_title_th":  staffResult.AccountTitleTh,
		"account_title_eng": staffResult.AccountTitleEng,
		"email":             staffResult.Email,
		"role":              staffResult.Role,
		"exp":               time.Now().Add(time.Hour * time.Duration(exp)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	return token.SignedString([]byte(os.Getenv("env")))
}
