package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matmust/pairStoring/mock"
	"github.com/matmust/pairStoring/pkg/flushing"
)

func TestFlushHandle(t *testing.T) {
	r := mock.PairRepository{}
	r.FlushFn = func() {}
	fs := flushing.NewService(&r)

	req, err := http.NewRequest("DELETE", "/pairs/", nil)
	if err != nil {
		t.Fatal(err)
	}
	s := NewServer(nil, nil, fs)
	rr := httptest.NewRecorder()

	s.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
