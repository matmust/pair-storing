// Package inmem provides in-memory implementations of all the domain repositories.
package inmem

import (
	"sync"

	"github.com/matmust/pairStoring"
)

type pairRepository struct {
	pairs sync.Map
}

// NewPairRepository returns a new instance of a in-memory pair repository.
func NewPairRepository() pairStoring.PairRepository {
	pairs := sync.Map{}
	return &pairRepository{pairs: pairs}
}

func (r *pairRepository) Set(p *pairStoring.Pair) bool {
	_, ok := r.pairs.Load(p.Key)
	r.pairs.Store(p.Key, p.Value)
	return ok
}

func (r *pairRepository) Get(key string) (string, error) {
	if value, ok := r.pairs.Load(key); ok {
		return value.(string), nil
	}
	return "", pairStoring.ErrKeyNotFound
}

func (r *pairRepository) Flush() {
	r.pairs.Range(func(key, value interface{}) bool {
		r.pairs.Delete(key)
		return true
	})
}

func (r *pairRepository) SetAll(m map[string]string) {
	for k, v := range m {
		r.pairs.Store(k, v)
	}
}

func (r *pairRepository) GetAll() map[string]string {
	m := make(map[string]string)

	r.pairs.Range(func(k, v interface{}) bool {
		m[k.(string)] = v.(string)
		return true
	})
	return m
}
