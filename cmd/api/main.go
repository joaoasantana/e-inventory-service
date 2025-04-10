package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joaoasantana/e-inventory-service/internal/api/router"
	"github.com/joaoasantana/e-inventory-service/internal/configs"
	"github.com/joaoasantana/e-inventory-service/internal/domain/usecase"
	"github.com/joaoasantana/e-inventory-service/internal/infra/repository"
	"github.com/joaoasantana/e-inventory-service/pkg/conn"
	"go.uber.org/zap"
)

func main() {
	config := configs.LoadAllConfig()

	logger := conn.DebugLogger()
	dbConn := conn.SQLDatabase(config.Database)

	defer closeVariables(dbConn, logger)

	logger = logger.With(
		zap.Any("app", config.App),
		zap.Any("server", config.Server),
	)

	// Dependency Injection
	categoryRepo := repository.NewCategoryRepository(dbConn)
	categoryUseCase := usecase.NewCategoryUseCase(logger, categoryRepo)

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		router.RegisterCategory(api, categoryUseCase)
	}

	if err := r.Run(config.Server.Port); err != nil {
		closeVariables(dbConn, logger)
		panic(err)
	}
}

func closeVariables(dbConn *sqlx.DB, logger *zap.Logger) {
	if err := dbConn.Close(); err != nil {
		panic(err)
	}

	if err := logger.Sync(); err != nil {
		panic(err)
	}
}
