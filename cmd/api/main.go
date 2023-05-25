package cmd

import (
	"log"
	"net/http"

	"github.com/alemelomeza/cautious-journey/internal/app"
	"github.com/alemelomeza/cautious-journey/internal/infra/db"
	"github.com/alemelomeza/cautious-journey/internal/infra/rest"
	"github.com/gorilla/mux"
)

func main() {
	repo := db.NewSqliteDB("demo.db")
	srv := app.NewChildService(repo, repo)
	hdl := rest.NewChildHandler(srv)

	r := mux.NewRouter()
	r.HandleFunc("/childrens", hdl.List).Methods(http.MethodGet)
	r.HandleFunc("/child/{name}", hdl.Get).Methods(http.MethodPost)

	s := &http.Server{
		Handler: r,
		Addr:    ":8000",
	}
	log.Fatal(s.ListenAndServe())
}
