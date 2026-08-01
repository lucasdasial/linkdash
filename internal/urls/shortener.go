package urls

import (
	"fmt"
	"lucasdasial/rupi/pkg/hashid"
	"math/rand/v2"
	"net/url"
)

type Service struct {
	encoder *hashid.Service
}

func New(salt string) (*Service, error) {
	encoder, err := hashid.New(salt)
	if err != nil {
		return nil, fmt.Errorf("shortener: %w", err)
	}

	return &Service{encoder: encoder}, nil
}

func (s *Service) Encode(rawURL string) (string, error) {
	if err := validateURL(rawURL); err != nil {
		return "", err
	}

	id := rand.IntN(1_000_000_000)

	code, err := s.encoder.Encode(id)
	if err != nil {
		return "", fmt.Errorf("shortener: encode id %d: %w", id, err)
	}

	return code, nil
}

func (s *Service) Decode(hash string) (int, error) {
	return s.encoder.Decode(hash)
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
