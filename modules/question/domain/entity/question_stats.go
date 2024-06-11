package entity

type QuestionStats struct {
	QuestionID    string `json:"questionId"`
	TotalAttempts int    `json:"totalAttempts"`
	CorrectAttempts  int    `json:"correctAttempts"`
}
