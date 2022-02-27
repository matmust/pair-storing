package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/matmust/pairStoring/pkg/flushing"
	"github.com/matmust/pairStoring/pkg/getting"
	"github.com/matmust/pairStoring/pkg/http/rest"
	"github.com/matmust/pairStoring/pkg/inmem"
	"github.com/matmust/pairStoring/pkg/setting"
	"github.com/matmust/pairStoring/pkg/storage"
)

const (
	defultFilePath = "./tmp/data.json"
	defaultPort    = "9000"
)

func main() {

	var (
		addr     = envString("PORT", defaultPort)
		httpAddr = flag.String("http.addr", "127.0.0.1:"+addr, "HTTP listen address")
	)
	flag.Parse()

	r := inmem.NewPairRepository()
	fileStorage := storage.NewFileStorage(defultFilePath, r)

	fileStorage.Load()
	fileStorage.PeriodicBackup(5 * time.Second)

	ss := setting.NewService(r)
	gs := getting.NewService(r)
	fls := flushing.NewService(r)

	s := rest.NewServer(ss, gs, fls)

	log.Printf("Server is starting up at http://127.0.0.1:%d/", 9000)

	err := http.ListenAndServe(*httpAddr, logRequest(s))

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
