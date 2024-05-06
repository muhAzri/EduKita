// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"EduKita/modules/learning_topics/data/repositories"
	"EduKita/modules/learning_topics/domain/usecases"
	"EduKita/modules/learning_topics/handler"
	"database/sql"
	"github.com/go-playground/validator"
	"github.com/google/wire"
)

// Injectors from learning_topics_handler.go:

func InitializeLearningTopicsHandler(db *sql.DB) *handler.LearningTopicHandler {
	learningTopicRepositoryImpl := repositories.NewLearningTopicRepository(db)
	getAllLearningTopicsUsecaseImpl := usecases.NewGetAllLearningTopicsUsecase(learningTopicRepositoryImpl)
	learningTopicHandler := handler.NewLearningTopicHandler(getAllLearningTopicsUsecaseImpl)
	return learningTopicHandler
}

// learning_topics_handler.go:

var LearningTopicsHandlerSet = wire.NewSet(repositories.NewLearningTopicRepository, wire.Bind(new(repositories.LearningTopicRepository), new(*repositories.LearningTopicRepositoryImpl)), usecases.NewGetAllLearningTopicsUsecase, wire.Bind(new(usecases.GetAllLearningTopicsUsecase), new(*usecases.GetAllLearningTopicsUsecaseImpl)), validator.New, handler.NewLearningTopicHandler)