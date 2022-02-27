package setting

import (
	"testing"

	"github.com/matmust/pairStoring/mock"
)

func TestSet(t *testing.T) {

	var r mock.PairRepository

	s := NewService(&r)

	r.GetFn = func(key string) (string, error) {
		return "test", nil
	}

	s.Set("test", "test")

	value, err := r.Get("test")
	if err != nil {
		t.Fatal(err)
	}

	if value != "test" {
		t.Errorf("value = %s; want = %s", value, "test")
	}

}
