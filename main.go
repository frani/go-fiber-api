package main

import (
	"github.com/frani/go-fiber-api/configs"

	"github.com/frani/go-fiber-api/middlewares"

	"github.com/frani/go-fiber-api/routes"

	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Connected with database
	configs.ConnectDB()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(limiter.New())
	app.Use(cors.New())
	app.Use(compress.New())

	// Create a /api/v1 endpoint
	api := app.Group("/")

	// Bind routes
	api.Use("/status", routes.StatusRoutes)
	api.Use("/users", routes.UsersRoutes)

	// Handle not founds
	app.Use("*", middlewares.NotFound)

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000
}
