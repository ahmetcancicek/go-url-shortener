package main

import (
	handler2 "github.com/ahmetcancicek/go-url-shortener/internal/app/shortener/handler"
	"github.com/ahmetcancicek/go-url-shortener/internal/app/shortener/repository/mongo"
	service2 "github.com/ahmetcancicek/go-url-shortener/internal/app/shortener/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {

	// DB Connect
	repo, err := mongo.NewRepository("mongodb://localhost:27017", "url-shortener", "admin", "password", 10)
	if err != nil {
		log.Fatal(err)
	}

	// Repository and service layer
	service := service2.NewRedirectService(repo)
	handler := handler2.NewHandler(service)

	// Router
	h := mux.NewRouter()
	h.HandleFunc("/api/v1/redirect", handler.CreateRedirect()).Methods(http.MethodPost)
	h.HandleFunc("/api/v1/redirect/{code}", handler.FindRedirectByCode()).Methods(http.MethodGet)

	// Server
	srv := &http.Server{
		Addr:         ":8500",
		ReadTimeout:  10 * time.Duration(time.Second),
		WriteTimeout: 10 * time.Duration(time.Second),
		Handler:      h,
	}

	// Start Server
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
