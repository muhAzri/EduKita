package model

import "github.com/google/uuid"

type UserProfileModel struct {
	ID                uuid.UUID        `json:"id"`
	Name              string           `json:"name"`
	Email             string           `json:"email"`
	ProfilePicture    string           `json:"profilePicture"`
	UserEloHistory    []UserEloHistory `json:"userEloHistory"`
	CurrentElo        int              `json:"currentElo"`
	TotalQuizAnswered int              `json:"totalQuizAnswered"`
	CreatedAt         int64            `json:"createdAt"`
	UpdatedAt         int64            `json:"updatedAt"`
}

type UserEloHistory struct {
	Elo       int    `json:"elo"`
	Date      string `json:"date"`
	UpdatedAt int64  `json:"-"`
}
