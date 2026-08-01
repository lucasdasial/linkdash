package urls

import "time"

type UrlModel struct {
	ID        int64     `db:"id" json:"-"`
	Code      string    `db:"code" json:"code"`
	RawURL    string    `db:"original_url" json:"raw_url"`
	Clicks    int64     `db:"clicks" json:"clicks"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// PORQUE UM PONTEIRO?
	ExpiresAt *time.Time `db:"expires_at" json:"expires_at,omitempty"`
}
