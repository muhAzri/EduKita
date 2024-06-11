package handler

import (
	"EduKita/modules/core/response"
	"EduKita/modules/question/data/model"
	"EduKita/modules/question/domain/usecases"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

type AdminQuestionHandler struct {
	AdminQuestionUsecase usecases.AdminQuestionUsecase
	Validate             *validator.Validate
}

func NewAdminQuestionHandler(AdminQuestionUsecase usecases.AdminQuestionUsecase, validate *validator.Validate) *AdminQuestionHandler {
	return &AdminQuestionHandler{
		AdminQuestionUsecase: AdminQuestionUsecase,
		Validate:             validate,
	}
}

func (h *AdminQuestionHandler) GetQuestion(w http.ResponseWriter, r *http.Request) {

	questionId := chi.URLParam(r, "question_id")

	question, err := h.AdminQuestionUsecase.GetQuestion(questionId)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code == "22P02" {
				response.BuildResponseFailure(http.StatusNotFound, "Cannot Find Question From That ID", w)
				return
			}
		}
		response.BuildResponseFailure(http.StatusInternalServerError, err.Error(), w)
		return
	}

	response.BuildResponseSuccess(http.StatusOK, "Get question success", "success", question, w)
}

func (h *AdminQuestionHandler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {

	questionId := chi.URLParam(r, "question_id")

	err := h.AdminQuestionUsecase.DeleteQuestion(questionId)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code == "22P02" {
				response.BuildResponseFailure(http.StatusNotFound, "Cannot Find Question From That ID", w)
				return
			}
		}
		response.BuildResponseFailure(http.StatusInternalServerError, err.Error(), w)
		return
	}

	response.BuildResponseSuccess(http.StatusOK, "Delete question success", "success", nil, w)
}

func (h *AdminQuestionHandler) UpdateQuestion(w http.ResponseWriter, r *http.Request) {

	var questionModel model.QuestionModel

	err := json.NewDecoder(r.Body).Decode(&questionModel)

	if err != nil {
		response.BuildResponseFailure(http.StatusBadRequest, "Invalid request body", w)
		return
	}

	err = h.Validate.Struct(questionModel)

	if err != nil {
		response.BuildResponseFailure(http.StatusBadRequest, err.Error(), w)
		return
	}

	err = h.AdminQuestionUsecase.UpdateQuestion(questionModel)

	if err != nil {
		response.BuildResponseFailure(http.StatusInternalServerError, err.Error(), w)
		return
	}

	response.BuildResponseSuccess(http.StatusOK, "Update question success", "success", nil, w)
}

func (h *AdminQuestionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {

	var questionModel model.QuestionModel

	err := json.NewDecoder(r.Body).Decode(&questionModel)

	if err != nil {
		response.BuildResponseFailure(http.StatusBadRequest, "Invalid request body", w)
		return
	}

	err = h.Validate.Struct(questionModel)

	if err != nil {
		response.BuildResponseFailure(http.StatusBadRequest, err.Error(), w)
		return
	}

	err = h.AdminQuestionUsecase.CreateQuestion(questionModel)

	if err != nil {
		response.BuildResponseFailure(http.StatusInternalServerError, err.Error(), w)
		return
	}

	response.BuildResponseSuccess(http.StatusOK, "Create question success", "success", nil, w)
}
