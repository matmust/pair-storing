package rest

import (
	"net/http"

	"github.com/matmust/pairStoring/pkg/flushing"
)

type flushingHandler struct {
	fs flushing.Service
}

// FlushHandle - Flush Pairs
// URL : /pairs/
// Method: DELETE
// Output: http status ok.
func (fh *flushingHandler) FlushHandle(w http.ResponseWriter, r *http.Request) {
	fh.fs.Flush()
	w.WriteHeader(http.StatusOK)

}
