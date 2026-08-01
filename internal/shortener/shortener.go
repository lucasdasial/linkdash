package shortener

import (
	"fmt"
	"math/rand/v2"
	"net/url"

	"lucasdasial/rupi/internal/base62"
	"lucasdasial/rupi/internal/hashid"
)

type Shortener struct {
	encoder *hashid.Service
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

	// TODO: id virá do incremento automático do banco de dados; por ora é aleatório.
	id := rand.IntN(1_000_000_000)

	code, err := s.encoder.Encode(int(base62.Decode(base62.Encode(id))))
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
