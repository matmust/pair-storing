package flushing

import (
	"testing"

	"github.com/matmust/pairStoring"
	"github.com/matmust/pairStoring/mock"
)

func TestFlush(t *testing.T) {
	var r mock.PairRepository

	s := NewService(&r)
	p := pairStoring.NewPair("test", "test")

	r.SetFn = func(p *pairStoring.Pair) bool {
		return true
	}

	r.Set(p)

	r.GetAllFn = func() map[string]string {
		return map[string]string{}
	}

	r.FlushFn = func() {}

	s.Flush()

	m := r.GetAll()

	if len(m) != 0 {
		t.Errorf("value = %v; want = %v", len(m), 0)
	}

}
