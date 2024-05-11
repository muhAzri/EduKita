package handler

import (
	"EduKita/modules/core/constants"
	"EduKita/modules/core/response"
	"EduKita/modules/question/data/model"
	"EduKita/modules/question/domain/usecases"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type AnswerHandler struct {
	AnswerQuestionUsecase usecases.AnswerQuestionUsecase
	Validate              *validator.Validate
}

func NewAnswerHandler(AnswerQuestionUsecase usecases.AnswerQuestionUsecase, validate *validator.Validate) *AnswerHandler {
	return &AnswerHandler{
		AnswerQuestionUsecase: AnswerQuestionUsecase,
		Validate:              validate,
	}
}

func (ah *AnswerHandler) AnswerQuestion(w http.ResponseWriter, r *http.Request) {

	var answerQuestionModel model.AnswerModel

	err := json.NewDecoder(r.Body).Decode(&answerQuestionModel)

	if err != nil {
		response.BuildResponseFailure(http.StatusBadRequest, "Invalid request body", w)
		return
	}

	err = ah.Validate.Struct(answerQuestionModel)

	if err != nil {
		response.BuildResponseFailure(http.StatusBadRequest, err.Error(), w)
		return
	}

	answeredResponse, err := ah.AnswerQuestionUsecase.AnswerQuestion(answerQuestionModel.QuestionId, *answerQuestionModel.AnswerIndex, r.Context().Value(constants.UserIDKey).(string))

	if err != nil {
		if err.Error() == "question already answered and cannot be reanswered again" {
			response.BuildResponseFailure(http.StatusBadRequest, err.Error(), w)
			return
		}

		response.BuildResponseFailure(http.StatusInternalServerError, err.Error(), w)
		return
	}

	response.BuildResponseSuccess(http.StatusOK, "Answer success", "success", answeredResponse, w)
}
