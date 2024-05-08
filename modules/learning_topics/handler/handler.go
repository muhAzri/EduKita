package handler

import (
	"EduKita/modules/core/response"
	"EduKita/modules/learning_topics/domain/usecases"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type LearningTopicHandler struct {
	GetAllLearningTopicsUsecase usecases.GetAllLearningTopicsUsecase
	Validate                    *validator.Validate
}

func NewLearningTopicHandler(getAllLearningTopicsUsecase usecases.GetAllLearningTopicsUsecase) *LearningTopicHandler {
	return &LearningTopicHandler{GetAllLearningTopicsUsecase: getAllLearningTopicsUsecase}
}

func (h *LearningTopicHandler) GetAllLearningTopics(w http.ResponseWriter, r *http.Request) {

	learningTopics, err := h.GetAllLearningTopicsUsecase.GetAllLearningTopics()
	if err != nil {
		response.BuildResponseFailure(http.StatusInternalServerError, err.Error(), w)
		return
	}

	response.BuildResponseSuccess(http.StatusOK, "Get all learning topics success", "success", learningTopics, w)
}
