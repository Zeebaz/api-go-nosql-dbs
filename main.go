package main

import (
	"github.com/Zeebaz/api-go-nosql-dbs/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// func handleAddPrediction(c *fiber.Ctx) error {
// 	match := Match{}
// 	if err := c.BodyParser(&match); err != nil {
// 		return err
// 	}
// 	// match.Id = uuid.NewString()

// 	return c.Status(fiber.StatusOK).JSON(match)
// }

// func handleGetAllPredictions(c *fiber.Ctx) error {
// 	rediscontroller.SetNewMatch("matches", "partido1")
// 	return c.SendString("All")
// }

func mainPoint(c *fiber.Ctx) error {
	return c.SendString("Hola mundo")
}

func main() {
	app := fiber.New()

	// middlewares
	app.Use(logger.New())
	app.Use(requestid.New())

	app.Get("/", mainPoint)
	app.Get("/matches", mainPoint)
	app.Post("/prediction", handlers.HandleAddNewPrediction)
	app.Get("/predictions", handlers.HandleGetPredictions)

	app.Listen(":3000")
}
