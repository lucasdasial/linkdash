-- +goose Up
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    code VARCHAR(20) NOT NULL UNIQUE,
    original_url TEXT NOT NULL,
    clicks BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ
);

CREATE UNIQUE INDEX idx_urls_code ON urls (code);

-- +goose Down
SELECT 'down SQL query';
DROP TABLE urls