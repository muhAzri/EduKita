//go:build wireinject
// +build wireinject

package wire

import (
	"EduKita/modules/profile/data/repositories"
	"EduKita/modules/profile/domain/usecases"
	"EduKita/modules/profile/handler"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

func NewValidator() *validator.Validate {
	return validator.New()
}

var ProfileHandlerSet = wire.NewSet(
	repositories.NewProfileRepository,
	wire.Bind(new(repositories.ProfileRepository), new(*repositories.ProfileRepositoryImpl)),
	usecases.NewGetProfileUsecase,
	wire.Bind(new(usecases.GetProfileUsecase), new(*usecases.GetProfileUsecaseImpl)),
	NewValidator, // Provide the validator options here
	handler.NewProfileHandler,
)

func InitializeProfileHandler(db *sql.DB) *handler.ProfileHandler {
	wire.Build(ProfileHandlerSet)
	return &handler.ProfileHandler{}
}
