package conn

import (
	"github.com/jmoiron/sqlx"
	"github.com/joaoasantana/e-inventory-service/pkg/config"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

func SQLDatabase(logger *zap.Logger, config config.DatabaseInfo) *sqlx.DB {
	logger.Info("connect database", zap.String("status", "connecting"))

	connection, err := sqlx.Open(config.Driver, config.URL())
	if err != nil {
		logger.Error("connect database", zap.String("status", "error"), zap.Error(err))

		panic(err)
	}

	if err = connection.Ping(); err != nil {
		logger.Error("connect database", zap.String("status", "error"), zap.Error(err))

		panic(err)
	}

	logger.Info("connect database", zap.String("status", "connected"))

	return connection
}
