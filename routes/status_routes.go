package routes

import (
	controller "github.com/frani/go-fiber-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func StatusRoutes(app *fiber.App) {
	// Create routes group.
	route := app.Group("/")
	route.Get("/", controller.GetStatus)
}
