package imagesrv

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ServeHTTP(r *mux.Router) {
	r.PathPrefix(cfg.URI).Handler(
		http.StripPrefix(
			cfg.URI, http.HandlerFunc(Image),
		),
	)
}
