package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"dododev/api/configs"
	"dododev/api/models"
)

func GetAllUsers(c *fiber.Ctx) error {
	collection := configs.GlobalMongoInstance.Db.Collection("users")

	// Define options for the Find method
	findOptions := options.Find()

	// Find all documents in the collection
	cursor, err := collection.Find(c.Context(), bson.D{}, findOptions)
	if err != nil {
		return c.Status(500).SendString("Failed to retrieve users from the database")
	}
	defer cursor.Close(c.Context())

	// Decode the documents into a slice of UserModel
	var users []models.UserModel
	if err := cursor.All(c.Context(), &users); err != nil {
		return c.Status(500).SendString("Failed to decode users from the database")
	}

	return c.JSON(users)
}
