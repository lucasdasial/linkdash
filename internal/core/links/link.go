package links

import "time"

type Link struct {
	ID        int       `json:"id"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
