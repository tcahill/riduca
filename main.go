package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	_ "github.com/lib/pq"
)

var (
	postgresURL string
	db          *sql.DB

	domainName string
)

func init() {
	flag.StringVar(&postgresURL,
		"postgres-url",
		"postgres://postgres:postgres@localhost/riduca?sslmode=disable",
		"URL for postgres DB")
	flag.StringVar(&domainName,
		"domain-name",
		"http://localhost:5000",
		"Domain name for this service")
}

func main() {
	var err error
	db, err = sql.Open("postgres", postgresURL)
	if err != nil {
		fmt.Printf("error establishing postgres connection: %s", err)
		os.Exit(1)
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/{id:[A-Za-z0-9]{6}}", expandURL)
	r.Post("/shorten", shortenURL)

	http.ListenAndServe(":5000", r)
}
