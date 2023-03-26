package routes

import (
	handlers "github.com/frani/go-fiber-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func StatusRoutes(app *fiber.App) {
	// Create routes group.
	route := app.Group("/status")
	route.Get("/", handlers.GetStatus)
}
