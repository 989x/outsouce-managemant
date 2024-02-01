package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"osm/api/database"
	"osm/api/helpers"
	"osm/api/models"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(c *fiber.Ctx) error {

	var login_body models.Login_body
	if err := c.BodyParser(&login_body); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	godotenv.Load(".env")
	validate := validator.New()
	if err := validate.Struct(login_body); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
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
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	buffer := bytes.NewBuffer(json_conv)

	client := &http.Client{}

	http_request, err := http.NewRequest("POST", "https://one.th/api/oauth/getpwd", buffer)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	http_request.Header.Set("Content-Type", "application/json")

	request_client, err := client.Do(http_request)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	defer request_client.Body.Close()

	http_result, err := ioutil.ReadAll(request_client.Body)

	var login_reqponse models.Login_response

	if err = json.Unmarshal(http_result, &login_reqponse); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	if login_reqponse.Result != "Success" {
		return helpers.JsonResponse(c, nil, 400, nil, "Fail : Username or Passwprd invalid.")
	}

	http_request, err = http.NewRequest("GET", "https://one.th/api/account", nil)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	http_request.Header.Set("Authorization", login_reqponse.TokenType+" "+login_reqponse.AccessToken)

	request_client, err = client.Do(http_request)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}
	defer request_client.Body.Close()

	http_result, err = ioutil.ReadAll(request_client.Body)
	var getOneAccount models.Get_OneAccount
	if err = json.Unmarshal(http_result, &getOneAccount); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	if err != nil {
		return helpers.JsonResponse(c, err, 503, nil, "Fail")
	}

	collection := database.MgConn.Db.Collection("staffs")
	context := database.MgConn.Ctx

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
		"fname_eng":         getOneAccount.FirstNameTh,
		"lname_eng":         getOneAccount.LastNameTh,
		"fname_th":          getOneAccount.FirstNameEng,
		"lname_th":          getOneAccount.LastNameEng,
		"account_title_th":  getOneAccount.SpecialTitleNameTh,
		"account_title_eng": getOneAccount.SpecialTitleNameEng,
		"email":             getOneAccount.Email,
		"role":              "Admin",
		"exp":               time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	return token.SignedString([]byte(os.Getenv("env")))
}
