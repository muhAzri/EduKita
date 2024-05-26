package repositories

import (
	"EduKita/modules/question/data/model"
	"EduKita/modules/question/domain/entity"
	"database/sql"
	"errors"
	"math"
	"time"

	"github.com/google/uuid"
)

type AnswerRepository interface {
	AnswerQuestion(QuestionID string, AnswerIndex int, UserId string) (model.AnsweredModel, error)
}

type AnswerRepositoryImpl struct {
	db *sql.DB
}

func NewAnswerRepository(db *sql.DB) *AnswerRepositoryImpl {
	return &AnswerRepositoryImpl{db: db}
}

func (r *AnswerRepositoryImpl) AnswerQuestion(QuestionID string, AnswerIndex int, UserId string) (model.AnsweredModel, error) {
	timeNow := time.Now().UTC().UnixMilli()

	tx, err := r.db.Begin()
	if err != nil {
		return model.AnsweredModel{}, err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var question entity.Question
	var questionStats entity.QuestionStats
	var isCorrect bool
	var latestScore int
	var additionalScore int
	var newScore int

	query := `SELECT is_correct FROM history_answers WHERE question_id = $1 AND user_id = $2 ORDER BY created_at DESC LIMIT 1`

	row := tx.QueryRow(query, QuestionID, UserId)

	var isAnsweredCorrect bool
	err = row.Scan(&isAnsweredCorrect)

	if err == nil {
		if isAnsweredCorrect {
			return model.AnsweredModel{}, errors.New("question already answered and cannot be reanswered again")
		}
	} else if err != sql.ErrNoRows {
		return model.AnsweredModel{}, err
	}

	query = `SELECT id, learning_topic_id, content, correct_answer_index, created_at, updated_at FROM questions WHERE id = $1`

	row = tx.QueryRow(query, QuestionID)

	err = row.Scan(&question.ID, &question.LearningTopicID, &question.Content, &question.CorrectAnswerIndex, &question.CreatedAt, &question.UpdatedAt)

	if err != nil {
		return model.AnsweredModel{}, err
	}

	query = `SELECT question_id,total_attempts, correct_attempts FROM question_stats WHERE question_id = $1`

	row = tx.QueryRow(query, question.ID)

	err = row.Scan(&questionStats.QuestionID, &questionStats.TotalAttempts, &questionStats.CorrectAttempts)

	if err != nil {
		return model.AnsweredModel{}, err
	}

	if question.CorrectAnswerIndex == AnswerIndex {
		isCorrect = true
		additionalScore = int(math.Round(CalculateQuestionDifficulty(questionStats.CorrectAttempts, questionStats.TotalAttempts) * 50))
		if additionalScore > 50 {
			additionalScore = 50
		} else if additionalScore < -50 {
			additionalScore = -50
		}

		questionStats.CorrectAttempts = questionStats.CorrectAttempts + 1
		questionStats.TotalAttempts = questionStats.TotalAttempts + 1
	} else {
		isCorrect = false
		additionalScore = int(math.Round(CalculateQuestionDifficulty(questionStats.CorrectAttempts, questionStats.TotalAttempts) * -50))

		if additionalScore > 50 {
			additionalScore = 50
		} else if additionalScore < -50 {
			additionalScore = -50
		}

		questionStats.TotalAttempts = questionStats.TotalAttempts + 1
	}

	query = `INSERT INTO history_answers (id, user_id, question_id, answer, is_correct, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err = tx.Exec(query, uuid.New().String(), UserId, question.ID, AnswerIndex, isCorrect, timeNow, timeNow)

	if err != nil {
		return model.AnsweredModel{}, err
	}

	query = `UPDATE question_stats SET total_attempts = $1, correct_attempts = $2 WHERE question_id = $3`

	_, err = tx.Exec(query, questionStats.TotalAttempts, questionStats.CorrectAttempts, question.ID)

	if err != nil {
		return model.AnsweredModel{}, err
	}

	query = `SELECT coalesce(new_elo, 1000) FROM user_elo_history WHERE user_id = $1 ORDER BY created_at DESC LIMIT 1`

	row = tx.QueryRow(query, UserId)

	err = row.Scan(&latestScore)

	if err != nil {
		if err == sql.ErrNoRows {
			latestScore = 1000
		} else {
			return model.AnsweredModel{}, err
		}

	}

	newScore = EloCalculator(latestScore, additionalScore)

	query = `INSERT INTO user_elo_history (id, user_id, previous_elo, new_elo, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = tx.Exec(query, uuid.New().String(), UserId, latestScore, newScore, timeNow, timeNow)

	if err != nil {
		return model.AnsweredModel{}, err
	}

	return model.AnsweredModel{IsCorrect: isCorrect, PreviousElo: latestScore, NewElo: newScore}, nil
}

func EloCalculator(previousScore, kFactor int) int {
	return previousScore + kFactor
}

func CalculateQuestionDifficulty(correctAnswersCount, totalUsers int) float64 {
	if totalUsers == 0 {
		return 0
	}
	if correctAnswersCount == 0 {
		return 1
	}

	return float64(correctAnswersCount) / float64(totalUsers)
}
