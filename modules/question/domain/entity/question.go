package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Question struct {
	ID                 uuid.UUID      `json:"id"`
	LearningTopicID    uuid.UUID      `json:"learningTopicId"`
	Content            string         `json:"content"`
	Answers            pq.StringArray `json:"answers"`
	CorrectAnswerIndex int            `json:"correctAnswerIndex"`
	CreatedAt          int64         `json:"createdAt"`
	UpdatedAt          int64         `json:"updatedAt"`
}
