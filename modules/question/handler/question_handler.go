package handler

import (
	"EduKita/modules/core/response"
	"EduKita/modules/question/domain/usecases"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

type QuestionHandler struct {
	GetQuestionByLearningTopicUsecase usecases.GetQuestionByLearningTopicUsecase
	Validate                          *validator.Validate
}

func NewQuestionHandler(getQuestionByLearningTopicUsecase usecases.GetQuestionByLearningTopicUsecase, validate *validator.Validate) *QuestionHandler {
	return &QuestionHandler{
		GetQuestionByLearningTopicUsecase: getQuestionByLearningTopicUsecase,
		Validate:                          validate,
	}
}

func (qh *QuestionHandler) GetQuestionByLearningTopic(w http.ResponseWriter, r *http.Request) {

	learningTopicID := chi.URLParam(r, "learning_topic_id")

	questions, err := qh.GetQuestionByLearningTopicUsecase.GetQuestionByLearningTopic(learningTopicID)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code == "22P02" {
				response.BuildResponseFailure(http.StatusNotFound, "Cannot Find Learning Topic From That ID", w)
				return
			}
		}
		response.BuildResponseFailure(http.StatusInternalServerError, err.Error(), w)
		return
	}

	response.BuildResponseSuccess(http.StatusOK, "Get question success", "success", questions, w)
}

func (qh *QuestionHandler) Get10RandomQuestionByLearningTopic(w http.ResponseWriter, r *http.Request) {

	learningTopicID := chi.URLParam(r, "learning_topic_id")

	questions, err := qh.GetQuestionByLearningTopicUsecase.Get10RandomQuestionByLearningTopic(learningTopicID)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code == "22P02" {
				response.BuildResponseFailure(http.StatusNotFound, "Cannot Find Learning Topic From That ID", w)
				return
			}
		}
		response.BuildResponseFailure(http.StatusInternalServerError, err.Error(), w)
		return
	}

	response.BuildResponseSuccess(http.StatusOK, "Get question success", "success", questions, w)
}
