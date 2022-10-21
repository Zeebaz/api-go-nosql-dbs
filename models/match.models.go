package models

type Match struct {
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	Phase int    `json:"phase"`
	Score string `json:"score"`
}

type Matches []Match
