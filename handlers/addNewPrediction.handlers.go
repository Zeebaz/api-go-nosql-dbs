package handlers

import (
	"fmt"

	"github.com/Zeebaz/api-go-nosql-dbs/models"
	"github.com/gofiber/fiber/v2"
)

func HandleAddNewPrediction(c *fiber.Ctx) error {
	newPredictionMatch := models.Match{}
	if err := c.BodyParser(&newPredictionMatch); err != nil {
		return c.SendStatus(fiber.StatusConflict)
	}
	prediction := fmt.Sprintf("%s:%s:%d:%s", newPredictionMatch.Team1, newPredictionMatch.Team2, newPredictionMatch.Phase, newPredictionMatch.Score)
	fmt.Println(prediction)
	return c.SendString("ff")
}
