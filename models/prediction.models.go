package models

type Prediction struct {
	Match string `json:"match"`
	Times int    `json:"times"`
}

type Predictions []Prediction
