package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cdxy1/go-file-storage/internal/entity"
	"github.com/cdxy1/go-file-storage/internal/storage/postgres"
)

type Metadata struct {
	Db *pgxpool.Pool
}

func NewMetadataRepo() (*Metadata, error) {
	db, err := postgres.NewPostgres()
	if err != nil {
		return nil, err
	}

	return &Metadata{Db: db}, nil
}

func (md *Metadata) Create(ctx context.Context, u *entity.Metadata) error {
	tx, err := md.Db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, `INSERT INTO "metadata"(id, name, url, size, type, created_at) VALUES ($1, $2, $3, $4, $5, $6)`, u.Id, u.Name, u.Url, u.Size, u.Type, u.CreatedAt); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (md *Metadata) GetByID(ctx context.Context, id string) (*entity.Metadata, error) {
	tx, err := md.Db.Begin(ctx)
	var f entity.Metadata
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, `SELECT id, name, url, size, type, created_at FROM "metadata" WHERE id=$1`, id)

	err = row.Scan(&f.Id, &f.Name, &f.Url, &f.Size, &f.Type, &f.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &f, nil
}

func (md *Metadata) GetAll(ctx context.Context) ([]entity.Metadata, error) {
	tx, err := md.Db.Begin(ctx)
	var fSl []entity.Metadata
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(ctx, `SELECT id, name, url, size, type, created_at FROM "metadata"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var f entity.Metadata

		if err := rows.Scan(&f.Id, &f.Name, &f.Url, &f.Size, &f.Type, &f.CreatedAt); err != nil {
			return nil, err
		}

		fSl = append(fSl, f)
	}
	return fSl, nil
}
