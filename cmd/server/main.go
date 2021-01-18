package main

import (
	"github.com/ahmetcancicek/go-url-shortener/config"
	handler2 "github.com/ahmetcancicek/go-url-shortener/internal/app/shortener/handler"
	"github.com/ahmetcancicek/go-url-shortener/internal/app/shortener/repository/mongo"
	service2 "github.com/ahmetcancicek/go-url-shortener/internal/app/shortener/service"
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
	service := service2.NewRedirectService(repo)
	handler := handler2.NewHandler(service)

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
