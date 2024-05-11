package model

type AnswerModel struct {
	QuestionId  string `json:"questionId" validate:"required"`
	AnswerIndex *int    `json:"answerIndex" validate:"required"`
}
