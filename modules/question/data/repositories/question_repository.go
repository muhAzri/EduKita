package repositories

import (
	"EduKita/modules/question/data/model"
	"EduKita/modules/question/domain/entity"
	"database/sql"

	"github.com/lib/pq"
)

type QuestionRepository interface {
	CreateQuestion(question entity.Question) (model.QuestionModel, error)
	GetQuestionByLearningTopic(learningTopicId string) ([]model.QuestionModel, error)
	Get10RandomQuestionByLearningTopic(userId, learningTopicId string) ([]model.QuestionModel, error)
	GetQuestionByID(id string) (model.QuestionModel, error)
	UpdateQuestion(question entity.Question) error
	DeleteQuestion(id string) error
}

type QuestionRepositoryImpl struct {
	db *sql.DB
}

func NewQuestionRepository(db *sql.DB) *QuestionRepositoryImpl {
	return &QuestionRepositoryImpl{db: db}
}

func (r *QuestionRepositoryImpl) CreateQuestion(question entity.Question) (model.QuestionModel, error) {

	query := `INSERT INTO questions (id, learning_topic_id, content, answers, correct_answer_index, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, learning_topic_id, content, answers, created_at, updated_at`

	row := r.db.QueryRow(query, question.ID, question.LearningTopicID, question.Content, question.Answers, question.CorrectAnswerIndex, question.CreatedAt, question.UpdatedAt)

	var questionModel model.QuestionModel
	err := row.Scan(&questionModel.ID, &questionModel.LearningTopicID, &questionModel.Content, &questionModel.Answers, &questionModel.CreatedAt, &questionModel.UpdatedAt)
	if err != nil {
		return model.QuestionModel{}, err
	}

	statsQuery := `INSERT INTO question_stats (question_id, total_attempts, correct_attempts) VALUES ($1, $2, $3)`

	_, err = r.db.Exec(statsQuery, questionModel.ID, 0, 0)

	if err != nil {
		return model.QuestionModel{}, err
	}

	return questionModel, nil
}

func (r *QuestionRepositoryImpl) GetQuestionByLearningTopic(learningTopicId string) ([]model.QuestionModel, error) {

	query := `SELECT id, learning_topic_id, content, answers, created_at, updated_at FROM questions WHERE learning_topic_id = $1`

	rows, err := r.db.Query(query, learningTopicId)

	if err != nil {
		return []model.QuestionModel{}, err
	}

	var questions []model.QuestionModel

	for rows.Next() {
		var question model.QuestionModel
		var answersRaw pq.StringArray

		err := rows.Scan(&question.ID, &question.LearningTopicID, &question.Content, &answersRaw, &question.CreatedAt, &question.UpdatedAt)
		if err != nil {
			return []model.QuestionModel{}, err
		}
		question.Answers = answersRaw
		questions = append(questions, question)
	}

	return questions, nil
}

func (r *QuestionRepositoryImpl) Get10RandomQuestionByLearningTopic(userId, learningTopicId string) ([]model.QuestionModel, error) {
	query := `
        SELECT q.id, q.learning_topic_id, q.content, q.answers, q.created_at, q.updated_at
        FROM questions q
        WHERE q.learning_topic_id = $1
        AND NOT EXISTS (
            SELECT 1
            FROM history_answers ha
            WHERE ha.question_id = q.id
            AND ha.user_id = $2
            AND ha.is_correct = true
        )
        ORDER BY RANDOM()
        LIMIT 10
    `

	rows, err := r.db.Query(query, learningTopicId, userId)
	if err != nil {
		return []model.QuestionModel{}, err
	}

	var questions []model.QuestionModel
	for rows.Next() {
		var question model.QuestionModel
		var answersRaw pq.StringArray

		err := rows.Scan(&question.ID, &question.LearningTopicID, &question.Content, &answersRaw, &question.CreatedAt, &question.UpdatedAt)
		if err != nil {
			return []model.QuestionModel{}, err
		}
		question.Answers = answersRaw
		questions = append(questions, question)
	}

	return questions, nil
}

func (r *QuestionRepositoryImpl) GetQuestionByID(id string) (model.QuestionModel, error) {

	query := `SELECT id, learning_topic_id, content, answers, created_at, updated_at FROM questions WHERE id = $1`

	row := r.db.QueryRow(query, id)

	var question model.QuestionModel
	err := row.Scan(&question.ID, &question.LearningTopicID, &question.Content, &question.Answers, &question.CreatedAt, &question.UpdatedAt)

	if err != nil {
		return model.QuestionModel{}, err
	}

	return question, nil
}

func (r *QuestionRepositoryImpl) UpdateQuestion(question entity.Question) error {

	query := `UPDATE questions SET content = $1, answers = $2, correct_answer_index = $3, updated_at = $4 WHERE id = $5`

	_, err := r.db.Exec(query, question.Content, question.Answers, question.CorrectAnswerIndex, question.UpdatedAt, question.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *QuestionRepositoryImpl) DeleteQuestion(id string) error {

	query := `DELETE FROM questions WHERE id = $1`

	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	queryStats := `DELETE FROM question_stats WHERE question_id = $1`

	_, err = r.db.Exec(queryStats, id)

	if err != nil {
		return err
	}

	return nil
}
