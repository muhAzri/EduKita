package handler

import (
	"EduKita/modules/core/response"
	"EduKita/modules/leaderboard/domain/usecases"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type LeaderboardHandler struct {
	GetLeaderboardUsecase usecases.GetLeaderboardUsecase
	Validate              *validator.Validate
}

func NewLeaderboardHandler(getLeaderboardUsecase usecases.GetLeaderboardUsecase, validate *validator.Validate) *LeaderboardHandler {
	return &LeaderboardHandler{
		GetLeaderboardUsecase: getLeaderboardUsecase,
		Validate:              validate,
	}
}

func (lh *LeaderboardHandler) GetLeaderboard(w http.ResponseWriter, r *http.Request) {

	leaderboard, err := lh.GetLeaderboardUsecase.GetLeaderboard()
	if err != nil {
		response.BuildResponseFailure(http.StatusInternalServerError, err.Error(), w)
		return
	}

	response.BuildResponseSuccess(http.StatusOK, "Get leaderboard success", "success", leaderboard, w)
}
