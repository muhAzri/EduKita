//go:build wireinject
// +build wireinject

package wire

import (
	"EduKita/modules/question/data/repositories"
	"EduKita/modules/question/domain/usecases"
	"EduKita/modules/question/handler"
	"database/sql"

	"github.com/google/wire"
)

var AdminQuestionHandler = wire.NewSet(
	repositories.NewQuestionRepository,
	wire.Bind(new(repositories.QuestionRepository), new(*repositories.QuestionRepositoryImpl)),
	usecases.NewAdminQuestionUsecase,
	wire.Bind(new(usecases.AdminQuestionUsecase), new(*usecases.AdminQuestionUsecaseImpl)),
	NewValidator,
	handler.NewAdminQuestionHandler,
)

func InitializeAdminQuestionHandler(db *sql.DB) *handler.AdminQuestionHandler {
	wire.Build(AdminQuestionHandler)
	return &handler.AdminQuestionHandler{}
}
