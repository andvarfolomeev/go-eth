package intrastructure

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

func NewDB(dataSourceName string) *sql.DB {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		slog.Error("Failed to open psql connection ", slog.Any("error", err))
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		slog.Error("Failed to ping psql", slog.Any("error", err))
		os.Exit(1)
	}

	return db
}
