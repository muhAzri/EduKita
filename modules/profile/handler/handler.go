package handler

import (
	"EduKita/modules/core/constants"
	"EduKita/modules/core/response"
	"EduKita/modules/profile/domain/usecases"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ProfileHandler struct {
	GetUserProfileUsecase usecases.GetProfileUsecase
	Validate              *validator.Validate
}

func NewProfileHandler(getUserProfileUsecase usecases.GetProfileUsecase, Validate *validator.Validate) *ProfileHandler {
	return &ProfileHandler{
		GetUserProfileUsecase: getUserProfileUsecase,
		Validate:              Validate,
	}
}

func (ph *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(constants.UserIDKey).(string)
	user, err := ph.GetUserProfileUsecase.GetProfile(userID)

	if err != nil {
		response.BuildResponseFailure(http.StatusInternalServerError, err.Error(), w)
		return
	}

	response.BuildResponseSuccess(http.StatusOK, "Get profile success", "success", user, w)
}
