package handler

import (
	"EduKita/modules/auth/data/model"
	"EduKita/modules/auth/domain/usecases"
	"EduKita/modules/core/constants"
	"EduKita/modules/core/response"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

type AuthHandler struct {
	usecase  usecases.Usecase
	validate *validator.Validate
}

func NewAuthHandler(usecase usecases.Usecase, validate *validator.Validate) *AuthHandler {
	return &AuthHandler{
		usecase:  usecase,
		validate: validate,
	}
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginModel model.LoginModel
	err := json.NewDecoder(r.Body).Decode(&loginModel)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		response.BuildResponseFailure(http.StatusBadRequest, "Invalid request body", w)
		return
	}

	err = ah.validate.Struct(loginModel)
	if err != nil {
		response.BuildResponseFailure(http.StatusBadRequest, err.Error(), w)
		return
	}

	user, err := ah.usecase.Login(loginModel, r.Context().Value(constants.UserIDKey).(string))
	if err != nil {
		response.BuildResponseFailure(http.StatusInternalServerError, err.Error(), w)
		return
	}

	response.BuildResponseSuccess(http.StatusOK, "Login success", "success", user, w)
}
