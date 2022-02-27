package rest

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matmust/pairStoring"
	"github.com/matmust/pairStoring/mock"
	"github.com/matmust/pairStoring/pkg/setting"
)

func TestSetHandle_NonExistsKey(t *testing.T) {
	r := mock.PairRepository{}
	r.SetFn = func(c *pairStoring.Pair) bool {
		return false
	}

	ss := setting.NewService(&r)

	tests := []TestRequest{
		{"/pairs/test", "test", http.StatusCreated, ""},
	}

	for _, testCase := range tests {

		req, err := http.NewRequest("PUT", testCase.requestUrl, bytes.NewBuffer([]byte(testCase.requestBody)))
		if err != nil {
			t.Fatal(err)
		}
		s := NewServer(ss, nil, nil)
		rr := httptest.NewRecorder()

		s.ServeHTTP(rr, req)
		if status := rr.Code; status != testCase.expectedStatusCode {
			t.Errorf("handler returned wrong status code: got %v want %v", status, testCase.expectedStatusCode)
		}

	}
}

func TestSetHandle_ExistsKey(t *testing.T) {
	r := mock.PairRepository{}
	r.SetFn = func(c *pairStoring.Pair) bool {
		return true
	}

	ss := setting.NewService(&r)

	tests := []TestRequest{
		{"/pairs/test", "test", http.StatusOK, ""},
	}

	for _, testCase := range tests {

		req, err := http.NewRequest("PUT", testCase.requestUrl, bytes.NewBuffer([]byte(testCase.requestBody)))
		if err != nil {
			t.Fatal(err)
		}
		s := NewServer(ss, nil, nil)
		rr := httptest.NewRecorder()

		s.ServeHTTP(rr, req)
		if status := rr.Code; status != testCase.expectedStatusCode {
			t.Errorf("handler returned wrong status code: got %v want %v", status, testCase.expectedStatusCode)
		}

	}
}
