package handlers

import (
	"fmt"
	"strconv"

	mongocontroller "github.com/Zeebaz/api-go-nosql-dbs/mongoController"
	"github.com/gofiber/fiber/v2"
)

func HandleCountPredictions(c *fiber.Ctx) error {
	// GET ALL JSON PREDICTIONS - MONGO
	response, err := mongocontroller.GetCountDocuments("PREDICTIONS", "matches")
	if err != nil {
		fmt.Println(err)
		return fiber.NewError(fiber.StatusConflict, "Error getting count predictions")
	}

	return c.Status(fiber.StatusOK).SendString(strconv.Itoa(int(response)))
}
