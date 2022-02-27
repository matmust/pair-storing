package getting

import (
	"testing"

	"github.com/matmust/pairStoring"
	"github.com/matmust/pairStoring/mock"
)

func TestGet(t *testing.T) {
	var r mock.PairRepository

	s := NewService(&r)
	p := pairStoring.NewPair("test", "test")

	r.GetFn = func(key string) (string, error) {
		return "test", nil
	}

	r.SetFn = func(p *pairStoring.Pair) bool {
		return true
	}

	r.Set(p)
	value, err := s.Get("test")
	if err != nil {
		t.Fatal(err)
	}

	if value != "test" {
		t.Errorf("value = %s; want = %s", value, "test")
	}

}
