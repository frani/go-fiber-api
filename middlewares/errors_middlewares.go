package middlewares

import "github.com/gofiber/fiber/v2"

func NotFound(c *fiber.Ctx) error {
	return c.Status(404).JSON(fiber.Map{"success": false, "message": "not found", "data": nil, "errors": nil})
}
