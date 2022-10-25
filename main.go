package main

import (
	"github.com/Zeebaz/api-go-nosql-dbs/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func mainPoint(c *fiber.Ctx) error {
	return c.SendString("Usactar 2022 - Go API")
}

func main() {
	app := fiber.New()

	// middlewares
	app.Use(logger.New())
	app.Use(requestid.New())

	app.Get("/", mainPoint)
	app.Get("/matches", handlers.HandleGetMatches)
	app.Post("/prediction", handlers.HandleAddNewPrediction)
	app.Get("/predictions", handlers.HandleGetPredictions)
	app.Get("/test", handlers.HandleTestEndpoint)

	app.Listen(":3000")
}
