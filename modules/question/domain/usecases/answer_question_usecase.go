package usecases

import (
	"EduKita/modules/question/data/model"
	"EduKita/modules/question/data/repositories"
)

type AnswerQuestionUsecase interface {
	AnswerQuestion(QuestionId string, AnswerIndex int, UserId string) (model.AnsweredModel, error)
}

type AnswerQuestionUsecaseImpl struct {
	AnswerRepository repositories.AnswerRepository
}

func NewAnswerQuestionUsecase(AnswerRepository repositories.AnswerRepository) *AnswerQuestionUsecaseImpl {
	return &AnswerQuestionUsecaseImpl{AnswerRepository: AnswerRepository}
}

func (u *AnswerQuestionUsecaseImpl) AnswerQuestion(QuestionId string, AnswerIndex int, UserId string) (model.AnsweredModel, error) {
	return u.AnswerRepository.AnswerQuestion(QuestionId, AnswerIndex, UserId)
}
