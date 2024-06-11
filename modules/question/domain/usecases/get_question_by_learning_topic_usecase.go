package usecases

import (
	"EduKita/modules/question/data/model"
	"EduKita/modules/question/data/repositories"
)

type GetQuestionByLearningTopicUsecase interface {
	GetQuestionByLearningTopic(learningTopicId string) ([]model.QuestionModel, error)
	Get10RandomQuestionByLearningTopic(userId, learningTopicId string) ([]model.QuestionModel, error)
}

type GetQuestionByLearningTopicUsecaseImpl struct {
	repository repositories.QuestionRepository
}

func NewGetQuestionByLearningTopicUsecase(repository repositories.QuestionRepository) *GetQuestionByLearningTopicUsecaseImpl {
	return &GetQuestionByLearningTopicUsecaseImpl{repository: repository}
}

func (u *GetQuestionByLearningTopicUsecaseImpl) GetQuestionByLearningTopic(learningTopicId string) ([]model.QuestionModel, error) {

	questions, err := u.repository.GetQuestionByLearningTopic(learningTopicId)

	if err != nil {
		return []model.QuestionModel{}, err
	}

	return questions, nil
}

func (u *GetQuestionByLearningTopicUsecaseImpl) Get10RandomQuestionByLearningTopic(userId, learningTopicId string) ([]model.QuestionModel, error) {

	questions, err := u.repository.Get10RandomQuestionByLearningTopic(userId, learningTopicId)

	if err != nil {
		return []model.QuestionModel{}, err
	}

	return questions, nil
}
