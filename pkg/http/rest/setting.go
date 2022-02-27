package rest

import (
	"io"
	"net/http"
	"strings"

	"github.com/matmust/pairStoring/pkg/setting"
	"github.com/matmust/pairStoring/pkg/validator"
)

type settingHandler struct {
	ss setting.Service
}

// SetHandle - Set Pair Value By Key
// URL : /pairs/{key}
// Path Parameter: string key
// Body : string value
// Method: PUT
// Output: http status ok if key found else http status created.
func (sh *settingHandler) SetHandle(w http.ResponseWriter, r *http.Request) {

	key := strings.TrimPrefix(r.URL.Path, "/pairs/")
	if validator.IsEmpty(key) {
		http.Error(w, "Key query parameter is missing", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	value, err := io.ReadAll(r.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		http.Error(w, "value parsing error", http.StatusBadRequest)
		return
	}

	if validator.IsEmpty(string(value)) {
		http.Error(w, "value is missing", http.StatusBadRequest)
		return
	}

	ok := sh.ss.Set(key, string(value))
	if ok {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
