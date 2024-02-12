package controllers

import (
	"outsource-management/api/configs"
	"outsource-management/api/helpers"
	"outsource-management/api/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllSkill(c *fiber.Ctx) error {

	collection := configs.MgConn.Db.Collection("skills")
	context := configs.MgConn.Ctx

	query := bson.D{{}}
	queryResult, err := collection.Find(context, query)
	if err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	var skills []models.Skills

	if err := queryResult.All(context, &skills); err != nil {
		return helpers.JsonResponse(c, err, 400, nil, "Fail")
	}

	return helpers.JsonResponse(c, nil, 200, skills, "Success")
}
