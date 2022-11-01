package handlers

import (
	"fmt"

	mongocontroller "github.com/Zeebaz/api-go-nosql-dbs/mongoController"
	"github.com/gofiber/fiber/v2"
)

func HandleMongoRecords(c *fiber.Ctx) error {
	// GET ALL JSON PREDICTIONS - MONGO
	response, err := mongocontroller.GetManyDocuments("PREDICTIONS", "matches")
	if err != nil {
		fmt.Println(err)
		return fiber.NewError(fiber.StatusConflict, "Error getting matches")
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
