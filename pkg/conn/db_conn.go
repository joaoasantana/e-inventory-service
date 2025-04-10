package conn

import (
	"github.com/jmoiron/sqlx"
	"github.com/joaoasantana/e-inventory-service/pkg/config"
	_ "github.com/lib/pq"
)

func SQLDatabase(config config.DatabaseInfo) *sqlx.DB {
	connection, err := sqlx.Open(config.Driver, config.URL())
	if err != nil {
		panic(err)
	}

	if err = connection.Ping(); err != nil {
		panic(err)
	}

	return connection
}
