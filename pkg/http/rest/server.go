package rest

import (
	"net/http"
	"regexp"

	"github.com/matmust/pairStoring/pkg/flushing"
	"github.com/matmust/pairStoring/pkg/getting"
	"github.com/matmust/pairStoring/pkg/setting"
)

var (
	getPairRe   = regexp.MustCompile(`^\/pairs\/(\w+)$`)
	setPairRe   = regexp.MustCompile(`^\/pairs\/(\w+)$`)
	flushPairRe = regexp.MustCompile(`^\/pairs[\/]*$`)
)

// Server holds the dependencies for a HTTP server.
type Server struct {
	settingService  setting.Service
	gettingService  getting.Service
	flushingService flushing.Service
	mux             *http.ServeMux
}

// NewServer returns a new HTTP server.
func NewServer(ss setting.Service, gs getting.Service, fls flushing.Service) *Server {
	mux := http.NewServeMux()
	s := &Server{ss, gs, fls, mux}

	mux.HandleFunc("/pairs", s.pairsHandler)
	mux.HandleFunc("/pairs/", s.pairsHandler)

	s.mux = mux
	return s

}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) pairsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch {
	case r.Method == http.MethodGet && getPairRe.MatchString(r.URL.Path):
		h := gettingHandler{s.gettingService}
		h.GetHandle(w, r)
		return
	case r.Method == http.MethodPut && setPairRe.MatchString(r.URL.Path):
		h := settingHandler{s.settingService}
		h.SetHandle(w, r)
		return
	case r.Method == http.MethodDelete && flushPairRe.MatchString(r.URL.Path):
		h := flushingHandler{s.flushingService}
		h.FlushHandle(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
