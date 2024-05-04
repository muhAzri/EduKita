package usecases

import (
	"EduKita/modules/auth/data/model"
	"EduKita/modules/auth/data/repositories"
	"EduKita/modules/auth/domain/entity"
	"EduKita/modules/core/utils"
	"time"

	"github.com/google/uuid"
)

type Usecase interface {
	Login(LoginModel model.LoginModel, FirebaseId string) (model.UserModel, error)
}

type UsecaseImpl struct {
	repository repositories.UserRepository
}

func NewUsecase(repository repositories.UserRepository) *UsecaseImpl {
	return &UsecaseImpl{repository}
}

func (u *UsecaseImpl) Login(LoginModel model.LoginModel, FirebaseId string) (model.UserModel, error) {

	gettedUser, err := u.repository.GetUserByFirebaseID(FirebaseId)

	if err != nil {
		return model.UserModel{}, err
	}

	if gettedUser.ID == "" {
		newUserModel := model.UserModel{
			ID:             uuid.New().String(),
			FirebaseID:     FirebaseId,
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
