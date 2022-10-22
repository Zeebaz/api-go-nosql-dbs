package handlers

import (
	"strconv"

	"github.com/Zeebaz/api-go-nosql-dbs/models"
	"github.com/Zeebaz/api-go-nosql-dbs/rediscontroller"
	"github.com/gofiber/fiber/v2"
)

func HandleGetPredictions(c *fiber.Ctx) error {

	allPredictions, err := rediscontroller.GetHGETALL("predictions")
	if err != nil {
		return c.SendStatus(fiber.StatusConflict)
	}

	predictionsReply := make(models.Predictions, 0)

	for k, v := range allPredictions {
		times, _ := strconv.Atoi(v)

		prediction := models.Prediction{
			Match: k,
			Times: times,
		}

		predictionsReply = append(predictionsReply, prediction)
	}

	return c.Status(fiber.StatusOK).JSON(predictionsReply)

}
