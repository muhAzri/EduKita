package model

type QuestionModel struct {
	ID              string   `json:"id"`
	LearningTopicID string   `json:"learningTopicId"`
	Content         string   `json:"content"`
	Answers         []string `json:"answers"`
	CreatedAt       string   `json:"createdAt"`
	UpdatedAt       string   `json:"updatedAt"`
}
