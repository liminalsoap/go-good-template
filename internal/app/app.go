package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"todo/config"
	"todo/internal/delivery/http"
	"todo/internal/repository"
	"todo/internal/usecase"
	"todo/pkg/logger"
	"todo/pkg/postgres"
)

func Run(cfg *config.Config) {
	log := logger.NewLogger(cfg.Logger.LogLevel)
	log.Info("Start")

	pg, err := postgres.NewDb(cfg.Postgres.PostgresqlUrl)
	if err != nil {
		log.Fatalf("failed to connect db: %s", pg)
	}
	defer pg.Conn.Close(context.Background())

	var useCases usecase.UseCases

	taskUseCase := usecase.NewTaskUseCase(
		repository.NewTaskRepo(pg),
	)
	useCases.Task = taskUseCase

	handler := gin.Default()
	http.NewRouter(handler, log, useCases)
	log.Fatalf("error: ", handler.Run(cfg.Http.Port))
}
