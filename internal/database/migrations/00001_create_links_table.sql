-- +goose Up
CREATE TABLE links(
    id SERIAL PRIMARY KEY,
    url TEXT NOT NULL,
    create_at TIMESTAMPTZ DEFAULT NOW()
);
-- +goose Down
DROP TABLE links;