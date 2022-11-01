package handlers

import (
	"fmt"

	mongocontroller "github.com/Zeebaz/api-go-nosql-dbs/mongoController"
	"github.com/gofiber/fiber/v2"
)

func HandleTestEndpoint(c *fiber.Ctx) error {
	// GET ALL JSON PREDICTIONS - MONGO
	response, err := mongocontroller.GetManyDocuments("PREDICTIONS", "matches")
	if err != nil {
		fmt.Println(err)
		return fiber.NewError(fiber.StatusConflict, "Error on tests")
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
