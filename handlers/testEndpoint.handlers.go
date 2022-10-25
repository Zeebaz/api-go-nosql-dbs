package handlers

import (
	mongocontroller "github.com/Zeebaz/api-go-nosql-dbs/mongoController"
	"github.com/gofiber/fiber/v2"
)

func HandleTestEndpoint(c *fiber.Ctx) error {
	// GET ALL JSON PREDICTIONS - MONGO
	response, err := mongocontroller.GetManyDocuments("PREDICTIONS", "matches")
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "Error getting matches")
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
