package usecases

import (
	"EduKita/modules/auth/data/model"
	"EduKita/modules/auth/data/repositories"
)

type GetShortProfileUsecase interface {
	GetShortProfile(id string) (model.ShortProfileModel, error)
}

type GetShortProfileUsecaseImpl struct {
	repository repositories.UserRepository
}

func NewGetShortProfileUsecase(repository repositories.UserRepository) *GetShortProfileUsecaseImpl {
	return &GetShortProfileUsecaseImpl{repository}
}

func (u *GetShortProfileUsecaseImpl) GetShortProfile(id string) (model.ShortProfileModel, error) {

	gettedUser, err := u.repository.GetUserByID(id)

	if err != nil {
		return model.ShortProfileModel{}, err
	}

	return model.ShortProfileModel{Name: gettedUser.Name}, nil

}
