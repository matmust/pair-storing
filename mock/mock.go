package mock

import (
	"github.com/matmust/pairStoring"
)

// PairRepository is a mock cargo repository.
type PairRepository struct {
	SetFn      func(c *pairStoring.Pair) bool
	SetInvoked bool

	GetFn      func(key string) (string, error)
	GetInvoked bool

	FlushFn      func()
	FlushInvoked bool

	GetAllFn      func() map[string]string
	GetAllInvoked bool

	SetAllFn      func(m map[string]string)
	SetAllInvoked bool
}

// Set calls the StoreFn.
func (r *PairRepository) Set(p *pairStoring.Pair) bool {
	r.SetInvoked = true
	if r != nil && r.SetFn != nil {
		return r.SetFn(p)
	}
	return true
}

// Get calls the FindFn.
func (r *PairRepository) Get(key string) (string, error) {
	r.GetInvoked = true
	if r != nil && r.GetFn != nil {
		return r.GetFn(key)
	}
	return "", nil
}

// Flush calls the FindAllFn.
func (r *PairRepository) Flush() {
	r.FlushInvoked = true
	r.FlushFn()
}

// SetAll calls the SetAllFn.
func (r *PairRepository) SetAll(m map[string]string) {
	r.GetInvoked = true
	r.SetAllFn(m)
}

// GetAll calls the GetAllFn.
func (r *PairRepository) GetAll() map[string]string {
	r.GetAllInvoked = true
	if r != nil && r.GetAllFn != nil {
		return r.GetAllFn()
	}

	return map[string]string{}
}
