package entity

import "github.com/google/uuid"

type Question struct {
	ID                 uuid.UUID `json:"id"`
	LearningTopicID    uuid.UUID `json:"learningTopicId"`
	Content            string    `json:"content"`
	Answers            []string  `json:"answers"`
	CorrectAnswerIndex int       `json:"correctAnswerIndex"`
	CreatedAt          string    `json:"createdAt"`
	UpdatedAt          string    `json:"updatedAt"`
}
