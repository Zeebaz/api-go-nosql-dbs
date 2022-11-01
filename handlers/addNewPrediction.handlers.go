package handlers

import (
	"fmt"
	"regexp"

	"github.com/Zeebaz/api-go-nosql-dbs/models"
	mongocontroller "github.com/Zeebaz/api-go-nosql-dbs/mongoController"
	"github.com/Zeebaz/api-go-nosql-dbs/rediscontroller"
	"github.com/gofiber/fiber/v2"
)

func HandleAddNewPrediction(c *fiber.Ctx) error {
	newPredictionMatch := models.Match{}
	if err := c.BodyParser(&newPredictionMatch); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	regexpresion := regexp.MustCompile(`^\d+-\d+`)
	isMatch := regexpresion.Match([]byte(newPredictionMatch.Score))
	if !isMatch {
		return fiber.NewError(fiber.StatusConflict, "Bad result format")
	}

	matchToVerify := fmt.Sprintf("%s:%s:%d", newPredictionMatch.Team1, newPredictionMatch.Team2, newPredictionMatch.Phase)
	verifyReply := isCorrectMatch(matchToVerify)
	if !verifyReply {
		return fiber.NewError(fiber.StatusConflict, "The match does not exist")
	}

	newPrediction := fmt.Sprintf("%s:%s:%d:%s", newPredictionMatch.Team1, newPredictionMatch.Team2, newPredictionMatch.Phase, newPredictionMatch.Score)
	response, err := rediscontroller.SetHINCRBY("predictions", newPrediction, 1)
	if err != nil && response > 0 {
		return fiber.NewError(fiber.StatusConflict, "An error occurred when inserting the prediction on redis db")
	}

	errm := mongocontroller.AddOneDocument(newPredictionMatch, "PREDICTIONS", "matches")
	if errm != nil {
		fmt.Println(errm)
		return fiber.NewError(fiber.StatusConflict, "An error occurred when inserting the prediction on mongo db")
	}

	return c.Status(fiber.StatusOK).SendString("Prediction was inserted")
}

func isCorrectMatch(match_ string) bool {
	allMatches, err := rediscontroller.GetSMEMBERS("matches")
	if err != nil {
		return false
	}

	for _, match := range allMatches {
		if match == match_ {
			return true
		}
	}

	return false
}
