package db

import (
	"database/sql"
	"log/slog"
)

func MustConnectDB(driver, URI string, maxOpenConns int) *sql.DB {
	database, err := sql.Open(driver, URI)
	if err != nil {
		panic(err)
	}

	err = database.Ping()
	if err != nil {
		panic(err)
	}

	database.SetMaxOpenConns(maxOpenConns)
	slog.Info("database connected")

	return database
}
