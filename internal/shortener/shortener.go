package shortener

import (
	"fmt"
	"net/url"
	"sync/atomic"

	"lucasdasial/rupi/internal/hashid"
)

type Shortener struct {
	encoder *hashid.Service
	nextID  atomic.Int64
}

func New(salt string) (*Shortener, error) {
	encoder, err := hashid.New(salt)
	if err != nil {
		return nil, fmt.Errorf("shortener: %w", err)
	}

	return &Shortener{encoder: encoder}, nil
}

// Shorten validates rawURL and returns a short code for it.
func (s *Shortener) Shorten(rawURL string) (string, error) {
	if err := validateURL(rawURL); err != nil {
		return "", err
	}

	id := s.nextID.Add(1)

	code, err := s.encoder.Encode(int(id))
	if err != nil {
		return "", fmt.Errorf("shortener: encode id %d: %w", id, err)
	}

	return code, nil
}

func validateURL(rawURL string) error {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return fmt.Errorf("shortener: invalid url %q: %w", rawURL, err)
	}

	if u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("shortener: invalid url %q: missing scheme or host", rawURL)
	}

	return nil
}
