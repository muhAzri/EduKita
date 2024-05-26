package handler

import (
	"EduKita/modules/auth/data/model"
	"EduKita/modules/auth/domain/usecases"
	"EduKita/modules/core/constants"
	"EduKita/modules/core/response"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	loginUsecase           usecases.LoginUsecase
	getShortProfileUsecase usecases.GetShortProfileUsecase
	generateTokenUsecase   usecases.GenerateTokenUsecase
	refreshTokenUsecase    usecases.RefreshTokenUsecase
	validate               *validator.Validate
}

func NewAuthHandler(loginUsecase usecases.LoginUsecase, generateTokenUsecase usecases.GenerateTokenUsecase, refreshTokenUsecase usecases.RefreshTokenUsecase, getShortProfileUsecase usecases.GetShortProfileUsecase, validate *validator.Validate) *AuthHandler {
	return &AuthHandler{
		loginUsecase:           loginUsecase,
		getShortProfileUsecase: getShortProfileUsecase,
		generateTokenUsecase:   generateTokenUsecase,
		refreshTokenUsecase:    refreshTokenUsecase,
		validate:               validate,
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

	user, err := ah.loginUsecase.Login(loginModel, r.Context().Value(constants.FirebaseIDKey).(string))
	if err != nil {
		response.BuildResponseFailure(http.StatusInternalServerError, err.Error(), w)
		return
	}

	jwtToken, err := ah.generateTokenUsecase.GenerateJWT(user.ID)

	if err != nil {
		response.BuildResponseFailure(http.StatusInternalServerError, err.Error(), w)
		return
	}

	response.BuildResponseSuccess(http.StatusOK, "Login success", "success", jwtToken, w)
}

func (ah *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var refreshModel model.RefreshModel
	err := json.NewDecoder(r.Body).Decode(&refreshModel)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		response.BuildResponseFailure(http.StatusBadRequest, "Invalid request body", w)
		return
	}

	err = ah.validate.Struct(refreshModel)
	if err != nil {
		response.BuildResponseFailure(http.StatusBadRequest, err.Error(), w)
		return
	}

	newToken, err := ah.refreshTokenUsecase.RefreshAccessToken(refreshModel.RefreshToken)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(refreshModel.RefreshToken)
		response.BuildResponseFailure(http.StatusInternalServerError, err.Error(), w)
		return
	}

	response.BuildResponseSuccess(http.StatusOK, "Refresh success", "success", newToken, w)
}

func (ah *AuthHandler) GetShortProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(constants.UserIDKey).(string)
	user, err := ah.getShortProfileUsecase.GetShortProfile(userID)
	if err != nil {
		response.BuildResponseFailure(http.StatusInternalServerError, err.Error(), w)
		return
	}
	response.BuildResponseSuccess(http.StatusOK, "Get short profile success", "success", user, w)
}
