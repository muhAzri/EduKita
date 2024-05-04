//go:build wireinject
// +build wireinject

package wire

import (
	"EduKita/modules/auth/data/repositories"
	"EduKita/modules/auth/domain/usecases"
	"EduKita/modules/auth/handler"
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/google/wire"
)

var AuthHandlerSet = wire.NewSet(
	repositories.NewUserRepository,
	wire.Bind(new(repositories.UserRepository), new(*repositories.UserRepositoryImpl)),
	usecases.NewUsecase,
	wire.Bind(new(usecases.Usecase), new(*usecases.UsecaseImpl)),
	validator.New,
	handler.NewAuthHandler,
)

func InitializeAuthHandler(db *sql.DB) *handler.AuthHandler {
	wire.Build(AuthHandlerSet)
	return &handler.AuthHandler{}
}
