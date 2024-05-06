//go:build wireinject
// +build wireinject

package wire

import (
	"EduKita/modules/learning_topics/data/repositories"
	"EduKita/modules/learning_topics/domain/usecases"
	"EduKita/modules/learning_topics/handler"
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/google/wire"
)

var LearningTopicsHandlerSet = wire.NewSet(
	repositories.NewLearningTopicRepository,
	wire.Bind(new(repositories.LearningTopicRepository), new(*repositories.LearningTopicRepositoryImpl)),
	usecases.NewGetAllLearningTopicsUsecase,
	wire.Bind(new(usecases.GetAllLearningTopicsUsecase), new(*usecases.GetAllLearningTopicsUsecaseImpl)),
	validator.New,
	handler.NewLearningTopicHandler,
)

func InitializeLearningTopicsHandler(db *sql.DB) *handler.LearningTopicHandler {
	wire.Build(LearningTopicsHandlerSet)
	return &handler.LearningTopicHandler{}
}
