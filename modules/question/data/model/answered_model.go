package model

type AnsweredModel struct {
	IsCorrect   bool `json:"isCorrect"`
	PreviousElo int  `json:"previousElo"`
	NewElo      int  `json:"newElo"`
}
