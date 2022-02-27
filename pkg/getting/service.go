// Package getting provides the use-case for getting value by key.
package getting

import "github.com/matmust/pairStoring"

// Service is the interface that provides getting methods.
type Service interface {
	// Get returns the existing value for the key if present. Otherwise, it returns an error
	Get(key string) (string, error)
}

type service struct {
	r pairStoring.PairRepository
}

func NewService(r pairStoring.PairRepository) Service {
	return &service{r}
}

func (s *service) Get(key string) (string, error) {
	return s.r.Get(key)
}
