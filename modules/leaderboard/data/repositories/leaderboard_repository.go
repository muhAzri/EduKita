package repositories

import (
	"EduKita/modules/leaderboard/data/model"
	"database/sql"
)

type LeaderboardRepository interface {
	GetLeaderBoard() ([]model.LeaderboardModel, error)
}

type LeaderboardRepositoryImpl struct {
	db *sql.DB
}

func NewLeaderboardRepository(db *sql.DB) *LeaderboardRepositoryImpl {
	return &LeaderboardRepositoryImpl{db}
}

func (l *LeaderboardRepositoryImpl) GetLeaderBoard() ([]model.LeaderboardModel, error) {
	query := `
	SELECT
    u.id AS userId,
    u.name,
    u.email,
    u.profile_picture AS profilePicture,
    e.new_elo AS elo
FROM
    users u
JOIN (
    SELECT DISTINCT ON (user_id)
        user_id,
        new_elo,
        created_at
    FROM
        user_elo_history
    ORDER BY
        user_id, created_at DESC
) e ON u.id = e.user_id
ORDER BY
    e.new_elo DESC
LIMIT 50;

	`
	rows, err := l.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var leaderboard []model.LeaderboardModel
	for rows.Next() {
		var leader model.LeaderboardModel
		if err := rows.Scan(&leader.UserID, &leader.Name, &leader.Email, &leader.ProfilePicture, &leader.Elo); err != nil {
			return nil, err
		}
		leaderboard = append(leaderboard, leader)
	}

	return leaderboard, nil
}
