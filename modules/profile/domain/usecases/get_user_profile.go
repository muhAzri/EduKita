package usecases

import (
	"EduKita/modules/profile/data/model"
	"EduKita/modules/profile/data/repositories"
)

type GetProfileUsecase interface {
	GetProfile(id string) (model.UserProfileModel, error)
}

type GetProfileUsecaseImpl struct {
	repository repositories.ProfileRepository
}

func NewGetProfileUsecase(repository repositories.ProfileRepository) *GetProfileUsecaseImpl {
	return &GetProfileUsecaseImpl{repository}
}

func (u *GetProfileUsecaseImpl) GetProfile(id string) (model.UserProfileModel, error) {

	gettedUser, err := u.repository.GetUserFromUserId(id)

	if err != nil {
		return model.UserProfileModel{}, err
	}

	return gettedUser, nil

}
