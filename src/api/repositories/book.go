package book

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"dododev/api/configs"
	"dododev/api/helpers"
	"dododev/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateBook creates a new book in MongoDB
func CreateBook(c *fiber.Ctx) error {
	// Parse request body
	var newBook models.Book
	if err := c.BodyParser(&newBook); err != nil {
		return helpers.HandleError(c, fiber.StatusBadRequest, "Invalid request payload", err)
	}

	// Set default values or perform validation if needed

	// Assign a new ObjectID for the book
	newBook.ID = primitive.NewObjectID()

	// Connect to MongoDB
	collection := configs.MongoClient.Database("booksDB").Collection("books")

	// Insert the new book document
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, newBook)
	if err != nil {
		return helpers.HandleError(c, fiber.StatusInternalServerError, "Failed to create book", err)
	}

	return c.Status(fiber.StatusCreated).JSON(newBook)
}

// GetAllBooks retrieves all books from MongoDB
func GetAllBooks(c *fiber.Ctx) error {
	// Connect to MongoDB
	collection := configs.MongoClient.Database("booksDB").Collection("books")

	// Query all books
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return helpers.HandleError(c, fiber.StatusInternalServerError, "Failed to retrieve books", err)
	}
	defer cursor.Close(ctx)

	// Iterate through the result set
	var books []models.Book
	if err := cursor.All(ctx, &books); err != nil {
		return helpers.HandleError(c, fiber.StatusInternalServerError, "Failed to parse books", err)
	}

	return c.JSON(books)
}

// GetBookByID retrieves a single book by ID from MongoDB
func GetBookByID(c *fiber.Ctx) error {
	// Parse book ID from the request parameters
	bookID := c.Params("id")

	// Convert book ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return helpers.HandleError(c, fiber.StatusBadRequest, "Invalid book ID", err)
	}

	// Connect to MongoDB
	collection := configs.MongoClient.Database("booksDB").Collection("books")

	// Query the book by ID
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var book models.Book
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&book)
	if err != nil {
		return helpers.HandleError(c, fiber.StatusNotFound, "Book not found", err)
	}

	return c.JSON(book)
}

// DeleteBook deletes a single book by ID from MongoDB
func DeleteBook(c *fiber.Ctx) error {
	// Parse book ID from the request parameters
	bookID := c.Params("id")

	// Convert book ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return helpers.HandleError(c, fiber.StatusBadRequest, "Invalid book ID", err)
	}

	// Connect to MongoDB
	collection := configs.MongoClient.Database("booksDB").Collection("books")

	// Delete the book by ID
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return helpers.HandleError(c, fiber.StatusInternalServerError, "Failed to delete book", err)
	}

	if result.DeletedCount == 0 {
		return helpers.HandleError(c, fiber.StatusNotFound, "Book not found", nil)
	}

	return c.JSON(fiber.Map{"message": "Book deleted successfully"})
}
