package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"outsource-management/api/configs"
	"outsource-management/api/helpers"
	"outsource-management/api/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func OneLogin(c *fiber.Ctx) error {

	var loginBody models.Login_body
	if err := c.BodyParser(&loginBody); err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}

	godotenv.Load(".env")
	validate := validator.New()
	if err := validate.Struct(loginBody); err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}

	loginRequest := models.Login_request{
		GrantType:    "password",
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Username:     loginBody.Username,
		Password:     loginBody.Password,
	}

	jsonConv, err := json.Marshal(loginRequest)
	if err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}
	buffer := bytes.NewBuffer(jsonConv)

	client := &http.Client{}

	httpRequest, err := http.NewRequest("POST", "https://one.th/api/oauth/getpwd", buffer)
	if err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	requestClient, err := client.Do(httpRequest)
	if err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}
	defer requestClient.Body.Close()

	httpResult, err := ioutil.ReadAll(requestClient.Body)

	var loginReqponse models.Login_response

	if err = json.Unmarshal(httpResult, &loginReqponse); err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}

	if loginReqponse.Result != "Success" {
		return helpers.JsonResponse(c, errors.New("Username or Passwprd invalid."), 503, nil, "Fail ")
	}

	return helpers.JsonResponse(c, nil, 200, loginReqponse, "Success")
}

func Login(c *fiber.Ctx) error {

	var login_body models.Login_body
	if err := c.BodyParser(&login_body); err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}

	godotenv.Load(".env")
	validate := validator.New()
	if err := validate.Struct(login_body); err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}

	loginRequest := models.Login_request{
		GrantType:    "password",
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Username:     login_body.Username,
		Password:     login_body.Password,
	}

	jsonConv, err := json.Marshal(loginRequest)
	if err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}
	buffer := bytes.NewBuffer(jsonConv)

	client := &http.Client{}

	httpRequest, err := http.NewRequest("POST", "https://one.th/api/oauth/getpwd", buffer)
	if err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	requestClient, err := client.Do(httpRequest)
	if err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}
	defer requestClient.Body.Close()

	httpResult, err := ioutil.ReadAll(requestClient.Body)

	var loginReqponse models.Login_response

	if err = json.Unmarshal(httpResult, &loginReqponse); err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}

	if loginReqponse.Result != "Success" {
		return helpers.JsonResponse(c, nil, 503, nil, "Fail : Username or Passwprd invalid.")
	}

	httpRequest, err = http.NewRequest("GET", "https://one.th/api/account", nil)
	if err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}

	httpRequest.Header.Set("Authorization", loginReqponse.TokenType+" "+loginReqponse.AccessToken)

	requestClient, err = client.Do(httpRequest)
	if err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}
	defer requestClient.Body.Close()

	httpResult, err = ioutil.ReadAll(requestClient.Body)
	var getOneAccount models.Get_OneAccount
	if err = json.Unmarshal(httpResult, &getOneAccount); err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}

	if err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}

	collection := configs.MgConn.Db.Collection("staffs")
	context := configs.MgConn.Ctx

	var staff models.StaffGetForUpdate

	query := bson.D{{Key: "account_id", Value: getOneAccount.ID}}
	if err := collection.FindOne(context, query).Decode(&staff); err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}

	act_token, err := CreateToken(&staff, &getOneAccount, "JWT_SECRET", 24)
	response_login := models.Admin_login_response{
		AccountId:   getOneAccount.ID,
		AccessToken: act_token,
	}

	return helpers.JsonResponse(c, nil, 200, response_login, "Success")
}

func Params(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return helpers.JsonResponse(c, nil, 200, claims, "Success")
}

func CreateToken(staff *models.StaffGetForUpdate, getOneAccount *models.Get_OneAccount, env string, exp int) (string, error) {
	cliams := jwt.MapClaims{
		"id":                staff.ID,
		"account_id":        getOneAccount.ID,
		"fname_eng":         getOneAccount.FirstNameEng,
		"lname_eng":         getOneAccount.LastNameEng,
		"fname_th":          getOneAccount.FirstNameTh,
		"lname_th":          getOneAccount.LastNameTh,
		"account_title_th":  getOneAccount.SpecialTitleNameTh,
		"account_title_eng": getOneAccount.SpecialTitleNameEng,
		"email":             getOneAccount.Email,
		"role":              "Admin",
		"exp":               time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	return token.SignedString([]byte(os.Getenv("env")))
}
