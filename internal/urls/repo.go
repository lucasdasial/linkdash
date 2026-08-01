package urls

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"lucasdasial/rupi/internal/models"
)

var ErrNotFound = errors.New("repository: url not found")

type URLRepo struct {
	pool *pgxpool.Pool
}

func NewURLRepository(pool *pgxpool.Pool) *URLRepo {
	return &URLRepo{pool: pool}
}

func (r *URLRepo) Create(ctx context.Context, code, rawURL string) (*models.URL, error) {
	const query = `
		INSERT INTO urls (code, original_url)
		VALUES ($1, $2)
		RETURNING id, code, original_url, clicks, created_at, expires_at
	`

	rows, err := r.pool.Query(ctx, query, code, rawURL)
	if err != nil {
		return nil, fmt.Errorf("repository: create url: %w", err)
	}

	url, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.URL])
	if err != nil {
		return nil, fmt.Errorf("repository: create url: %w", err)
	}

	return &url, nil
}

func (r *URLRepo) GetByCode(ctx context.Context, code string) (*models.URL, error) {
	const query = `
		SELECT id, code, original_url, clicks, created_at, expires_at
		FROM urls
		WHERE code = $1
	`

	rows, err := r.pool.Query(ctx, query, code)
	if err != nil {
		return nil, fmt.Errorf("repository: get url: %w", err)
	}

	url, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.URL])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("repository: get url: %w", err)
	}

	return &url, nil
}

func (r *URLRepo) List(ctx context.Context, limit, offset int) ([]models.URL, error) {
	const query = `
		SELECT id, code, original_url, clicks, created_at, expires_at
		FROM urls
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("repository: list urls: %w", err)
	}

	urls, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.URL])
	if err != nil {
		return nil, fmt.Errorf("repository: list urls: %w", err)
	}

	return urls, nil
}
