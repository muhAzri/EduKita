// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"EduKita/modules/leaderboard/data/repositories"
	"EduKita/modules/leaderboard/domain/usecases"
	"EduKita/modules/leaderboard/handler"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

// Injectors from leaderboard_handler.go:

func InitializeLeaderboardHandler(db *sql.DB) *handler.LeaderboardHandler {
	leaderboardRepositoryImpl := repositories.NewLeaderboardRepository(db)
	getLeaderboardUsecaseImpl := usecases.NewGetLeaderboardUsecase(leaderboardRepositoryImpl)
	validate := NewValidator()
	leaderboardHandler := handler.NewLeaderboardHandler(getLeaderboardUsecaseImpl, validate)
	return leaderboardHandler
}

// leaderboard_handler.go:

func NewValidator() *validator.Validate {
	return validator.New()
}

var LeaderboardHandlerSet = wire.NewSet(repositories.NewLeaderboardRepository, wire.Bind(new(repositories.LeaderboardRepository), new(*repositories.LeaderboardRepositoryImpl)), usecases.NewGetLeaderboardUsecase, wire.Bind(new(usecases.GetLeaderboardUsecase), new(*usecases.GetLeaderboardUsecaseImpl)), NewValidator, handler.NewLeaderboardHandler)
