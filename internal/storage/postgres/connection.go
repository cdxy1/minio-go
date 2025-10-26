package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cdxy1/minio-go/internal/config"
)

func NewPostgres(log *slog.Logger) (*pgxpool.Pool, error) {
	cfg := config.GetConfig()
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Database,
	)

	log.Info("Connecting to PostgreSQL",
		"host", cfg.Postgres.Host,
		"port", cfg.Postgres.Port,
		"database", cfg.Postgres.Database,
	)

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Error("Failed to create PostgreSQL pool", "error", err)
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Error("PostgreSQL ping failed", "error", err)
		pool.Close()
		return nil, err
	}

	log.Info("PostgreSQL connection established successfully")

	_, err = pool.Exec(context.Background(), `CREATE TABLE metadata (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		url TEXT,
		size INT NOT NULL,
		type TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL
		);`)
	if err != nil {
		log.Error("Failed to create 'files' table or table already exists", "error", err)
	}

	log.Info("Table 'files' ensured in PostgreSQL successfully")
	return pool, nil
}
