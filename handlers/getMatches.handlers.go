package handlers

import (
	"strconv"
	"strings"

	"github.com/Zeebaz/api-go-nosql-dbs/models"
	"github.com/Zeebaz/api-go-nosql-dbs/rediscontroller"
	"github.com/gofiber/fiber/v2"
)

func HandleGetMatches(c *fiber.Ctx) error {
	allMatches, err := rediscontroller.GetSMEMBERS("matches")
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "Error getting matches")
	}

	matchesReply := make(models.Matches, 0)

	for _, match := range allMatches {
		rest := strings.Split(match, ":")
		numberPhase, _ := strconv.Atoi(rest[2])

		matchReply := models.Match{
			Team1: rest[0],
			Team2: rest[1],
			Phase: numberPhase,
			Score: "",
		}
		matchesReply = append(matchesReply, matchReply)
	}

	return c.Status(fiber.StatusOK).JSON(matchesReply)
}
