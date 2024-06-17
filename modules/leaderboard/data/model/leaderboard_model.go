package model

import "github.com/google/uuid"

type LeaderboardModel struct {
	UserID         uuid.UUID `json:"userId"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	ProfilePicture string    `json:"profilePicture"`
	Elo            int       `json:"elo"`
}
