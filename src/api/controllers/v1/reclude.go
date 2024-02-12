package controllers

import (
	"net/http"
	"outsource-management/api/configs"
	"outsource-management/api/helpers"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func MethodPost(c *fiber.Ctx) error {
	return helpers.JsonResponse(c, nil, 200, "Hello, RecludeMethod POST endpoint", "success")
}

func MethodPut(c *fiber.Ctx) error {
	return helpers.JsonResponse(c, nil, 200, "Hello, RecludeMethod PUT endpoint", "success")
}

func MethodDelete(c *fiber.Ctx) error {
	return helpers.JsonResponse(c, nil, 200, "Hello, RecludeMethod DELETE endpoint", "success")
}

func GetMonthlyEmployeeAttendance(c *fiber.Ctx) error {
	// Get query parameters
	yearFromStr := c.Query("year_from")
	yearToStr := c.Query("year_to")

	// Initialize start and end years to search
	startYear := 0
	endYear := time.Now().Year()

	// Convert query parameters to integers if they are not empty
	if yearFromStr != "" {
		yearFrom, err := strconv.Atoi(yearFromStr)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid 'year_from' parameter",
			})
		}
		startYear = yearFrom
	}

	if yearToStr != "" {
		yearTo, err := strconv.Atoi(yearToStr)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid 'year_to' parameter",
			})
		}
		endYear = yearTo
	}

	// MongoDB collection and context
	collection := configs.MgConn.Db.Collection("staffs")
	context := configs.MgConn.Ctx

	// Define filter for querying
	filter := bson.M{}

	// Add conditions to filter based on years
	if startYear > 0 && endYear > 0 {
		filter["start_date"] = bson.M{
			"$gte": time.Date(startYear, time.January, 1, 0, 0, 0, 0, time.UTC),
			"$lte": time.Date(endYear, time.December, 31, 23, 59, 59, 0, time.UTC),
		}
	}

	// Count total staffs
	totalStaffs, err := collection.CountDocuments(context, filter)
	if err != nil {
		return err
	}

	// Aggregation pipeline to count employees per month
	pipeline := bson.A{
		bson.M{
			"$match": filter,
		},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"$dateToString": bson.M{
						"format": "%Y-%m",
						"date":   "$start_date",
					},
				},
				"count": bson.M{"$sum": 1},
			},
		},
		bson.M{
			"$sort": bson.M{"_id": 1},
		},
	}

	// Perform aggregation
	cursor, err := collection.Aggregate(context, pipeline)
	if err != nil {
		return err
	}
	defer cursor.Close(context)

	// Store the results in a map
	monthlyCounts := make(map[string]int)
	for cursor.Next(context) {
		var result struct {
			Month string `bson:"_id"`
			Count int    `bson:"count"`
		}
		if err := cursor.Decode(&result); err != nil {
			return err
		}
		monthlyCounts[result.Month] = result.Count
	}

	response := map[string]interface{}{
		"total_staffs":   totalStaffs,
		"monthly_counts": monthlyCounts,
	}
	return helpers.JsonResponse(c, nil, 200, response, "Monthly Employee Attendance")
}
