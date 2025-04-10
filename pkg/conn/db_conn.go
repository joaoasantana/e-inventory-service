package conn

import (
	"database/sql"
	"github.com/joaoasantana/e-inventory-service/pkg/config"
)

func OpenSQLDatabase(config config.DatabaseInfo) *sql.DB {
	connection, err := sql.Open(config.Driver, config.URL())
	if err != nil {
		panic(err)
	}

	if err = connection.Ping(); err != nil {
		panic(err)
	}

	return connection
}

func CloseSQLDatabase(connection *sql.DB) {
	if err := connection.Close(); err != nil {
		panic(err)
	}
}
