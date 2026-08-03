package links

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	db *pgxpool.Pool
}

func NewLinkRepo(pool *pgxpool.Pool) *Repo {
	return &Repo{
		db: pool,
	}
}

func (repo *Repo) Create(ctx context.Context, url string) (*Link, error) {
	rows, err := repo.db.Query(ctx, `
	INSERT INTO links (url)
	VALUES($1)
	RETURNING id, url, created_at;
	`, url)

	if err != nil {
		return nil, err
	}

	link, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Link])

	if err != nil {
		return nil, err
	}

	return &link, nil
}
