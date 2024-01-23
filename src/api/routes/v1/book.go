package v1

import (
	book "dododev/api/repositories"

	"github.com/gofiber/fiber/v2"
)

// SetupBookRoutes sets up book-related routes
func SetupBookRoutes(app *fiber.App) {
	// Create a group for v1
	v1 := app.Group("/api/v1")

	// Routes for books
	v1.Post("/books", book.CreateBook)
	v1.Get("/books", book.GetAllBooks)
	v1.Get("/books/:id", book.GetBookByID)
	v1.Delete("/books/:id", book.DeleteBook)
}
