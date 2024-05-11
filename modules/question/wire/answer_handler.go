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

var AnswerHandler = wire.NewSet(
	repositories.NewAnswerRepository,
	wire.Bind(new(repositories.AnswerRepository), new(*repositories.AnswerRepositoryImpl)),
	usecases.NewAnswerQuestionUsecase,
	wire.Bind(new(usecases.AnswerQuestionUsecase), new(*usecases.AnswerQuestionUsecaseImpl)),
	NewValidator,
	handler.NewAnswerHandler,
)

func InitializeAnswerHandler(db *sql.DB) *handler.AnswerHandler {
	wire.Build(AnswerHandler)
	return &handler.AnswerHandler{}
}
