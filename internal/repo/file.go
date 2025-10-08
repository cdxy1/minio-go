package repo

import (
	"context"
	"github.com/cdxy1/go-file-storage/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type File struct {
	Db *pgxpool.Pool
}

func (fr *File) Create(ctx context.Context, u *entity.File) error {
	tx, err := fr.Db.Begin(ctx)

	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, `INSERT INTO "files"(name, url) VALUES ($1,$2)`, u.Name, u.Url); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (fr *File) GetByID(ctx context.Context, id int) (*entity.File, error) {
	tx, err := fr.Db.Begin(ctx)
	var f entity.File

	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, `SELECT id, name, url, created_at FROM "files" WHERE id=$1`, id)

	err = row.Scan(&f)

	if err != nil {
		return nil, err
	}

	return &f, nil
}
