package usecases

import (
	"EduKita/modules/question/data/model"
	"EduKita/modules/question/data/repositories"
	"EduKita/modules/question/domain/entity"
	"time"

	"github.com/google/uuid"
)

type AdminQuestionUsecase interface {
	GetQuestion(id string) (model.QuestionModel, error)
	CreateQuestion(question model.QuestionModel) error
	UpdateQuestion(question model.QuestionModel) error
	DeleteQuestion(id string) error
}

type AdminQuestionUsecaseImpl struct {
	repository repositories.QuestionRepository
}

func NewAdminQuestionUsecase(repository repositories.QuestionRepository) *AdminQuestionUsecaseImpl {
	return &AdminQuestionUsecaseImpl{repository: repository}
}

func (u *AdminQuestionUsecaseImpl) GetQuestion(id string) (model.QuestionModel, error) {

	question, err := u.repository.GetQuestionByID(id)

	if err != nil {
		return model.QuestionModel{}, err
	}

	return question, nil
}

func (u *AdminQuestionUsecaseImpl) CreateQuestion(question model.QuestionModel) error {

	entity := entity.Question{
		ID:                 uuid.New(),
		LearningTopicID:    uuid.MustParse(question.LearningTopicID),
		Content:            question.Content,
		Answers:            question.Answers,
		CorrectAnswerIndex: question.CorrectAnswerIndex,
		CreatedAt:          time.Now().UnixMilli(),
		UpdatedAt:          time.Now().UnixMilli(),
	}

	_, err := u.repository.CreateQuestion(entity)

	if err != nil {
		return err
	}

	return nil

}

func (u *AdminQuestionUsecaseImpl) UpdateQuestion(question model.QuestionModel) error {
	return u.repository.UpdateQuestion(entity.Question{
		ID:                 uuid.MustParse(question.ID),
		LearningTopicID:    uuid.MustParse(question.LearningTopicID),
		Content:            question.Content,
		Answers:            question.Answers,
		CorrectAnswerIndex: question.CorrectAnswerIndex,
		UpdatedAt:          question.UpdatedAt,
	})
}

func (u *AdminQuestionUsecaseImpl) DeleteQuestion(id string) error {

	return u.repository.DeleteQuestion(id)
}
