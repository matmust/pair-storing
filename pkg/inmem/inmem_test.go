package inmem

import (
	"sync"
	"testing"

	"github.com/matmust/pairStoring"
)

func TestSet_NonExistsKey(t *testing.T) {

	pr := NewPairRepository()

	type testCase struct {
		name string
		p    *pairStoring.Pair
	}

	testCases := []testCase{
		{
			name: "Set value by key",
			p: &pairStoring.Pair{Key: "test",
				Value: "test"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ok := pr.Set(tc.p)

			value, _ := pr.Get(tc.p.Key)
			if ok && value != tc.p.Value {
				t.Errorf("Expected value %v, got %v", tc.p.Value, value)
			}

		})
	}

}

func TestSet_ExistsKey(t *testing.T) {

	pr := NewPairRepository()
	pr.Set(pairStoring.NewPair("test", "test"))
	type testCase struct {
		name string
		p    *pairStoring.Pair
	}

	testCases := []testCase{
		{
			name: "Set value by exists key",
			p: &pairStoring.Pair{Key: "test",
				Value: "test1"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ok := pr.Set(tc.p)
			value, _ := pr.Get(tc.p.Key)
			if !ok && value != tc.p.Value {
				t.Errorf("Expected value %v, got %v", tc.p.Value, value)
			}

		})
	}

}

func TestGet(t *testing.T) {

	pairs := sync.Map{}
	pairs.Store("test", "test")
	pr := pairRepository{pairs}

	type testCase struct {
		name        string
		key         string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Get value by key",
			key:         "test",
			expectedErr: nil,
		}, {
			name:        "Get value by non-exists key",
			key:         "nonExistKey",
			expectedErr: pairStoring.ErrKeyNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := pr.Get(tc.key)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}
}

func TestFlush(t *testing.T) {
	pairs := sync.Map{}
	pairs.Store("test", "test")
	pr := pairRepository{pairs}

	pr.Flush()

	if len(pr.GetAll()) != 0 {
		t.Errorf("Expected 0 , got %d", len(pr.GetAll()))
	}
}

func TestGetAll(t *testing.T) {
	pairs := sync.Map{}
	pairs.Store("test", "test")
	pr := pairRepository{pairs}

	if len(pr.GetAll()) != 1 {
		t.Errorf("Expected 1 , got %d", len(pr.GetAll()))
	}
}

func TestSetAll(t *testing.T) {
	pr := NewPairRepository()
	m := map[string]string{"test": "test"}
	pr.SetAll(m)

	if len(pr.GetAll()) != 1 {
		t.Errorf("Expected 1 , got %d", len(pr.GetAll()))
	}
}
