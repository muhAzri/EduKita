//go:build wireinject
// +build wireinject

package wire

import (
	"EduKita/modules/question/data/repositories"
	"EduKita/modules/question/domain/usecases"
	"EduKita/modules/question/handler"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

var QuestionHandlerSet = wire.NewSet(
	repositories.NewQuestionRepository,
	wire.Bind(new(repositories.QuestionRepository), new(*repositories.QuestionRepositoryImpl)),
	usecases.NewGetQuestionByLearningTopicUsecase,
	wire.Bind(new(usecases.GetQuestionByLearningTopicUsecase), new(*usecases.GetQuestionByLearningTopicUsecaseImpl)),
	NewValidator,
	handler.NewQuestionHandler,
)

func InitializeQuestionHandler(db *sql.DB) *handler.QuestionHandler {
	wire.Build(QuestionHandlerSet)
	return &handler.QuestionHandler{}
}

func NewValidator() *validator.Validate {
	return validator.New()
}
