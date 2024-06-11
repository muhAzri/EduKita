package entity

import "github.com/google/uuid"

type LearningTopic struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Icon        string    `json:"icon"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}
