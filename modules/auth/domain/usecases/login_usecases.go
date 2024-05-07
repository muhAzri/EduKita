package usecases

import (
	"EduKita/modules/auth/data/model"
	"EduKita/modules/auth/data/repositories"
	"EduKita/modules/auth/domain/entity"
	"EduKita/modules/core/utils"
	"time"

	"github.com/google/uuid"
)

type LoginUsecase interface {
	Login(LoginModel model.LoginModel, FirebaseId string) (model.UserModel, error)
}

type LoginUsecaseImpl struct {
	repository repositories.UserRepository
}

func NewLoginUsecase(repository repositories.UserRepository) *LoginUsecaseImpl {
	return &LoginUsecaseImpl{repository}
}

func (u *LoginUsecaseImpl) Login(LoginModel model.LoginModel, FirebaseId string) (model.UserModel, error) {

	gettedUser, err := u.repository.GetUserByFirebaseId(FirebaseId)

	if err != nil {
		return model.UserModel{}, err
	}

	if gettedUser.ID == "" {
		newUserModel := model.UserModel{
			ID:             uuid.New().String(),
			FirebaseId:     FirebaseId,
			Name:           LoginModel.Name,
			Email:          LoginModel.Email,
			ProfilePicture: LoginModel.ProfilePicture,
			CreatedAt:      time.Now().UTC().UnixMilli(),
			UpdatedAt:      time.Now().UTC().UnixMilli(),
		}

		newEntity, err := utils.AnyToType[entity.User](newUserModel)

		if err != nil {
			return model.UserModel{}, err
		}

		createdUser, err := u.repository.CreateUser(newEntity)

		if err != nil {
			return model.UserModel{}, err
		}

		return createdUser, nil
	}

	return gettedUser, nil
}
