package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cdxy1/go-file-storage/internal/config"
)

func NewPostgres() (*pgxpool.Pool, error) {
	cfg := config.GetConfig()
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Database)

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, err
	}

	_, err = pool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTSfiles (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		url TEXT,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
	)`)

	if err != nil {
		return nil, err
	}

	return pool, nil
}
