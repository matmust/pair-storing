// Package flushing provides the use-case for deleting all the keys.
package flushing

import "github.com/matmust/pairStoring"

// Service is the interface that provides flushing methods.
type Service interface {
	// Flush deletes the values for all keys.
	Flush()
}

type service struct {
	r pairStoring.PairRepository
}

func NewService(r pairStoring.PairRepository) Service {
	return &service{r}
}

func (s *service) Flush() {
	s.r.Flush()
}
