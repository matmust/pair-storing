// Package setting provides the use-case for setting value by key.
package setting

import "github.com/matmust/pairStoring"

// Service is the interface that provides setting methods.
type Service interface {
	// Set replaces the existing value for the key if present. Otherwise, it stores value by key
	Set(key, value string) bool
}
type service struct {
	r pairStoring.PairRepository
}

func NewService(r pairStoring.PairRepository) Service {
	return &service{r}
}

func (s *service) Set(key, value string) bool {

	p := pairStoring.NewPair(key, value)

	return s.r.Set(p)
}
