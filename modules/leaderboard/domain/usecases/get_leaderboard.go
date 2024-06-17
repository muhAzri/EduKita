package usecases

import (
	"EduKita/modules/leaderboard/data/model"
	"EduKita/modules/leaderboard/data/repositories"
)

type GetLeaderboardUsecase interface {
	GetLeaderboard() ([]model.LeaderboardModel, error)
}

type GetLeaderboardUsecaseImpl struct {
	repository repositories.LeaderboardRepository
}

func NewGetLeaderboardUsecase(repository repositories.LeaderboardRepository) *GetLeaderboardUsecaseImpl {
	return &GetLeaderboardUsecaseImpl{repository}
}

func (u *GetLeaderboardUsecaseImpl) GetLeaderboard() ([]model.LeaderboardModel, error) {

	gettedLeaderboard, err := u.repository.GetLeaderBoard()

	if err != nil {
		return []model.LeaderboardModel{}, err
	}

	return gettedLeaderboard, nil
}
