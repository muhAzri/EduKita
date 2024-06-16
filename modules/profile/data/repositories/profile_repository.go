package repositories

import (
	"EduKita/modules/profile/data/model"
	"database/sql"
	"sort"
	"time"
)

type ProfileRepository interface {
	GetUserFromUserId(userId string) (model.UserProfileModel, error)
}

type ProfileRepositoryImpl struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepositoryImpl {
	return &ProfileRepositoryImpl{db}
}

func (pr *ProfileRepositoryImpl) GetUserFromUserId(userId string) (model.UserProfileModel, error) {
	var userProfile model.UserProfileModel

	query := `SELECT id, name, email, profile_picture, created_at, updated_at FROM users WHERE id = $1`
	row := pr.db.QueryRow(query, userId)

	err := row.Scan(&userProfile.ID, &userProfile.Name, &userProfile.Email, &userProfile.ProfilePicture, &userProfile.CreatedAt, &userProfile.UpdatedAt)
	if err != nil {
		return model.UserProfileModel{}, err
	}

	userEloHistory, err := pr.GetUserEloHistory(userId)
	if err != nil {
		return model.UserProfileModel{}, err
	}

	if len(userEloHistory) == 0 {
		userProfile.UserEloHistory = []model.UserEloHistory{}
	} else {
		userProfile.UserEloHistory = userEloHistory
	}

	totalQuizAnsweredQuery := `SELECT COUNT(*) FROM history_answers WHERE user_id = $1`
	totalQuizAnsweredRow := pr.db.QueryRow(totalQuizAnsweredQuery, userId)

	err = totalQuizAnsweredRow.Scan(&userProfile.TotalQuizAnswered)
	if err != nil {
		return model.UserProfileModel{}, err
	}

	currentEloQuery := `SELECT new_elo FROM user_elo_history WHERE user_id = $1 ORDER BY updated_at DESC LIMIT 1`
	currentEloRow := pr.db.QueryRow(currentEloQuery, userId)

	var currentElo int
	err = currentEloRow.Scan(&currentElo)
	if err != nil && err != sql.ErrNoRows {
		return model.UserProfileModel{}, err
	}

	if err == sql.ErrNoRows {
		currentElo = 1000
	}

	userProfile.CurrentElo = currentElo

	return userProfile, nil
}

func (pr *ProfileRepositoryImpl) GetUserEloHistory(userId string) ([]model.UserEloHistory, error) {
	var userEloHistory []model.UserEloHistory

	query := `SELECT new_elo, updated_at FROM user_elo_history WHERE user_id = $1 AND updated_at >= $2 ORDER BY updated_at ASC`
	sevenDaysAgo := time.Now().AddDate(0, 0, -7).UnixNano() / int64(time.Millisecond)
	rows, err := pr.db.Query(query, userId, sevenDaysAgo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	eloHistoryMap := make(map[string]model.UserEloHistory)

	for rows.Next() {
		var newElo int
		var updatedAt int64
		err := rows.Scan(&newElo, &updatedAt)
		if err != nil {
			return nil, err
		}

		date := time.Unix(0, updatedAt*int64(time.Millisecond)).Format("2006-01-02")
		if existing, exists := eloHistoryMap[date]; !exists || updatedAt > existing.UpdatedAt {
			eloHistoryMap[date] = model.UserEloHistory{
				Elo:       newElo,
				Date:      time.Unix(0, updatedAt*int64(time.Millisecond)).Format(time.RFC3339),
				UpdatedAt: updatedAt,
			}
		}
	}

	for _, eloHistory := range eloHistoryMap {
		userEloHistory = append(userEloHistory, eloHistory)
	}

	sort.Slice(userEloHistory, func(i, j int) bool {
		return userEloHistory[i].Date < userEloHistory[j].Date
	})

	return userEloHistory, nil
}
