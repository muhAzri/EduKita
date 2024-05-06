package usecases

import (
	"EduKita/modules/learning_topics/data/model"
	"EduKita/modules/learning_topics/data/repositories"
)

type GetAllLearningTopicsUsecase interface {
	GetAllLearningTopics() ([]model.LearningTopicModel, error)
}

type GetAllLearningTopicsUsecaseImpl struct {
	repository repositories.LearningTopicRepository
}

func NewGetAllLearningTopicsUsecase(repository repositories.LearningTopicRepository) *GetAllLearningTopicsUsecaseImpl {
	return &GetAllLearningTopicsUsecaseImpl{repository: repository}
}

func (u *GetAllLearningTopicsUsecaseImpl) GetAllLearningTopics() ([]model.LearningTopicModel, error) {
	learningTopics, err := u.repository.GetAllLearningTopics()

	if err != nil {
		return []model.LearningTopicModel{}, err
	}

	if len(learningTopics) == 0 {
		return []model.LearningTopicModel{}, nil
	}

	return learningTopics, nil
}
