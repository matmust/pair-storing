package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matmust/pairStoring/mock"
	"github.com/matmust/pairStoring/pkg/getting"
	"github.com/matmust/pairStoring/pkg/setting"
)

type TestRequest struct {
	requestUrl           string
	requestBody          string
	expectedStatusCode   int
	expectedResponseBody string
}

func TestGetHandle(t *testing.T) {
	r := mock.PairRepository{}
	r.GetFn = func(key string) (string, error) {
		return "test", nil
	}
	gs := getting.NewService(&r)
	ss := setting.NewService(&r)

	tests := []TestRequest{
		{"/pairs/test", "", http.StatusOK, "test"},
	}

	for _, testCase := range tests {

		req, err := http.NewRequest("GET", testCase.requestUrl, nil)

		if err != nil {
			t.Fatal(err)
		}
		s := NewServer(ss, gs, nil)
		rr := httptest.NewRecorder()

		s.ServeHTTP(rr, req)
		if status := rr.Code; status != testCase.expectedStatusCode {
			t.Errorf("handler returned wrong status code: got %v want %v", status, testCase.expectedStatusCode)
		}

		if rr.Body.String() != testCase.expectedResponseBody {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), testCase.expectedResponseBody)
		}

	}
}
