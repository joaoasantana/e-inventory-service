package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joaoasantana/e-inventory-service/internal/api/router"
	"github.com/joaoasantana/e-inventory-service/internal/configs"
	"github.com/joaoasantana/e-inventory-service/pkg/conn"
	"go.uber.org/zap"
)

func main() {
	config := configs.LoadAllConfig()

	logger := conn.DebugLogger().With(
		zap.Any("app", config.App),
		zap.Any("server", config.Server),
	)

	dbConn := conn.SQLDatabase(config.Database)

	defer closeVariables(dbConn, logger)

	r := gin.Default()

	api := r.Group("/api/v1")
	{

		router.InitCategoryRoute(dbConn, logger, api)
		router.InitProductRoute(dbConn, logger, api)
		router.InitSupplierRoute(dbConn, logger, api)
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
