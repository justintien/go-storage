package main

import (
	"filesrv"
	"imagesrv"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	// DB     *gorm.DB
}

func main() {
	a := App{}

	serve(a)
}

func serve(a App) {
	a.Router = mux.NewRouter().StrictSlash(true)
	a.Init()
	a.Run(":" + os.Getenv("GO_EXPOSED_PORT"))
}

func (a *App) Init() {
	a.initializeRoutes()
}

func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func (a *App) Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server alive!"))
	}).Methods("GET")

	filesrv.ServeHTTP(a.Router)
	imagesrv.ServeHTTP(a.Router)
	a.Router.Use(Logger)
}

func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
