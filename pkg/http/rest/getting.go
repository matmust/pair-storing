package rest

import (
	"net/http"
	"strings"

	"github.com/matmust/pairStoring/pkg/getting"
	"github.com/matmust/pairStoring/pkg/validator"
)

type gettingHandler struct {
	gs getting.Service
}

// GetHandle - Get Pair Value By Key
// URL : /pairs/{key}
// Path Parameter: string key
// Method: GET
// Output: Pair string value if found else http status not found.
func (gh *gettingHandler) GetHandle(w http.ResponseWriter, r *http.Request) {

	key := strings.TrimPrefix(r.URL.Path, "/pairs/")
	if validator.IsEmpty(key) {
		http.Error(w, "Key query parameter is missing", http.StatusBadRequest)
		return
	}

	v, f := gh.gs.Get(key)

	if f == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(v))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
