package repositories

import (
	"EduKita/modules/learning_topics/data/model"
	"EduKita/modules/learning_topics/domain/entity"
	"database/sql"
)

type LearningTopicRepository interface {
	CreateLearningTopic(topic entity.LearningTopic) (model.LearningTopicModel, error)
	GetAllLearningTopics() ([]model.LearningTopicModel, error)
	UpdateLearningTopic(topic entity.LearningTopic) error
	DeleteLearningTopic(id string) error
}

type LearningTopicRepositoryImpl struct {
	db *sql.DB
}

func NewLearningTopicRepository(db *sql.DB) *LearningTopicRepositoryImpl {
	return &LearningTopicRepositoryImpl{db: db}
}

func (r *LearningTopicRepositoryImpl) CreateLearningTopic(topic entity.LearningTopic) (model.LearningTopicModel, error) {

	query := `INSERT INTO learning_topics (name, slug, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, slug, description, created_at, updated_at`

	row := r.db.QueryRow(query, topic.Name, topic.Slug, topic.Description, topic.CreatedAt, topic.UpdatedAt)

	var learningTopic model.LearningTopicModel

	err := row.Scan(&learningTopic.ID, &learningTopic.Name, &learningTopic.Slug, &learningTopic.Description, &learningTopic.CreatedAt, &learningTopic.UpdatedAt)

	if err != nil {
		return model.LearningTopicModel{}, err
	}

	return learningTopic, nil
}

func (r *LearningTopicRepositoryImpl) GetAllLearningTopics() ([]model.LearningTopicModel, error) {

	query := `SELECT id, name, slug, description, created_at, updated_at FROM learning_topics`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	var learningTopics []model.LearningTopicModel

	for rows.Next() {
		var learningTopic model.LearningTopicModel
		err := rows.Scan(&learningTopic.ID, &learningTopic.Name, &learningTopic.Slug, &learningTopic.Description, &learningTopic.CreatedAt, &learningTopic.UpdatedAt)
		if err != nil {
			return nil, err
		}
		learningTopics = append(learningTopics, learningTopic)
	}

	return learningTopics, nil
}

func (r *LearningTopicRepositoryImpl) UpdateLearningTopic(topic entity.LearningTopic) error {

	query := `UPDATE learning_topics SET name = $1, slug = $2, description = $3, updated_at = $4 WHERE id = $5`

	_, err := r.db.Exec(query, topic.Name, topic.Slug, topic.Description, topic.UpdatedAt, topic.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *LearningTopicRepositoryImpl) DeleteLearningTopic(id string) error {

	query := `DELETE FROM learning_topics WHERE id = $1`

	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
