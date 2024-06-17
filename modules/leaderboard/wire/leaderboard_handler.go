//go:build wireinject
// +build wireinject

package wire

import (
	"EduKita/modules/leaderboard/data/repositories"
	"EduKita/modules/leaderboard/domain/usecases"
	"EduKita/modules/leaderboard/handler"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

func NewValidator() *validator.Validate {
	return validator.New()
}

var LeaderboardHandlerSet = wire.NewSet(
	repositories.NewLeaderboardRepository,
	wire.Bind(new(repositories.LeaderboardRepository), new(*repositories.LeaderboardRepositoryImpl)),
	usecases.NewGetLeaderboardUsecase,
	wire.Bind(new(usecases.GetLeaderboardUsecase), new(*usecases.GetLeaderboardUsecaseImpl)),
	NewValidator,
	handler.NewLeaderboardHandler,
)

func InitializeLeaderboardHandler(db *sql.DB) *handler.LeaderboardHandler {
	wire.Build(LeaderboardHandlerSet)
	return &handler.LeaderboardHandler{}
}
