package model

import "github.com/lib/pq"

type QuestionModel struct {
	ID                 string         `json:"id"`
	LearningTopicID    string         `json:"learningTopicId"`
	Content            string         `json:"content"`
	Answers            pq.StringArray `json:"answers"`
	CorrectAnswerIndex int            `json:"correctAnswerIndex,omitempty"`
	CreatedAt          int64            `json:"createdAt"`
	UpdatedAt          int64            `json:"updatedAt"`
}
