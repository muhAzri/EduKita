package model

type LearningTopicModel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Icon        string `json:"icon"`
	CreatedAt   int    `json:"createdAt"`
	UpdatedAt   int    `json:"updatedAt"`
}
