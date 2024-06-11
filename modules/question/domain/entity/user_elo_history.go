package entity

import "github.com/google/uuid"

type UserEloHistory struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"userId"`
	PreviousElo int       `json:"previousElo"`
	NewElo      int       `json:"newElo"`
	CreatedAt   int       `json:"createdAt"`
	UpdatedAt   int       `json:"updatedAt"`
}
