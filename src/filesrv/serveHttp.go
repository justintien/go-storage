package filesrv

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

func ServeHTTP(r *mux.Router) {
	r.HandleFunc(path.Join(cfg.URI, "upload"), Upload).Methods("POST")
	r.PathPrefix(cfg.URI).Handler(
		http.StripPrefix(
			cfg.URI, http.FileServer(http.Dir(cfg.Root)),
		),
	)
}
