package controllers

import "github.com/gofiber/fiber/v2"

// GetBooks func gets all exists books.
// @Description Get all exists books.
// @Summary get all exists books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /v1/books [get]

func GetStatus(c *fiber.Ctx) error {

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"success": true,
		"data":    nil,
		"message": "ok",
		"errors":  nil,
	})
}
