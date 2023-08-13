package routes

import (
	"github.com/frani/go-fiber-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func UsersRoutes(app *fiber.App) {
	// Create routes group.
	route := app.Group("/users")
	route.Get("/", handlers.ListUsers)
	route.Post("/", handlers.PostUser)
	route.Get("/:userId", handlers.GetUser)
	route.Patch("/:userId", handlers.PatchUser)
	route.Delete("/:userId", handlers.DeleteUser)
}
