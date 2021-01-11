package main

import (
	"github.com/ahmetcancicek/go-url-shortener/config"
	"github.com/ahmetcancicek/go-url-shortener/repository/mongo"
	"github.com/ahmetcancicek/go-url-shortener/shortener"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {

	// Config SetUp
	cfg, err := config.SetUp()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := mongo.NewRepository(
		cfg.Database.URI,
		cfg.Database.Name,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Timeout,
	)
	if err != nil {
		log.Fatal(err)
	}
	service := shortener.NewRedirectService(repo)
	handler := shortener.NewHandler(service)

	h := mux.NewRouter()
	h.HandleFunc("/api/v1/redirect", handler.CreateRedirect()).Methods(http.MethodPost)
	h.HandleFunc("/api/v1/redirect/{code}", handler.FindRedirectByCode()).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		ReadTimeout:  10 * time.Duration(time.Second),
		WriteTimeout: 10 * time.Duration(time.Second),
		Handler:      h,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
