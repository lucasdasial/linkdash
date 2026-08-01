package hashid

import (
	"errors"

	hashids "github.com/speps/go-hashids/v2"
)

type Service struct {
	hash *hashids.HashID
}

func New(salt string) (*Service, error) {
	data := hashids.NewData()
	data.Salt = salt
	data.MinLength = 4

	h, err := hashids.NewWithData(data)
	if err != nil {
		return nil, err
	}

	return &Service{
		hash: h,
	}, nil
}

func (s *Service) Encode(id int) (string, error) {
	return s.hash.Encode([]int{id})
}

func (s *Service) Decode(hash string) (int, error) {
	ids, err := s.hash.DecodeWithError(hash)
	if err != nil {
		return 0, err
	}

	if len(ids) == 0 {
		return 0, errors.New("invalid hashid")
	}

	return ids[0], nil
}
