package routes

import (
	controller "github.com/frani/go-fiber-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func UsersRoutes(app *fiber.App) {
	// Create routes group.
	route := app.Group("/")
	route.Get("/", controller.ListUsers)
	route.Post("/", controller.PostUser)
	route.Get("/:userId", controller.GetUser)
	route.Patch("/:userId", controller.PatchUser)
	route.Delete("/:userId", controller.DeleteUser)
}
