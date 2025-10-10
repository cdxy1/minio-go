package repo

import (
	"context"
	"github.com/cdxy1/go-file-storage/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Metadata struct {
	Db *pgxpool.Pool
}

func (md *Metadata) Create(ctx context.Context, u *entity.Metadata) error {
	tx, err := md.Db.Begin(ctx)

	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, `INSERT INTO "files"(name, url) VALUES ($1,$2)`, u.Name, u.Url); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (md *Metadata) GetByID(ctx context.Context, id int) (*entity.Metadata, error) {
	tx, err := md.Db.Begin(ctx)
	var f entity.Metadata

	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, `SELECT id, name, url, created_at FROM "files" WHERE id=$1`, id)

	err = row.Scan(&f.Id, &f.Name, &f.Url, &f.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &f, nil
}
