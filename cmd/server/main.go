package main

import (
	"context"
	"github.com/ahmetcancicek/go-url-shortener/internal/app/shortener/handler"
	"github.com/ahmetcancicek/go-url-shortener/internal/app/shortener/repository/mongo"
	"github.com/ahmetcancicek/go-url-shortener/internal/app/shortener/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {

	// Context
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(2*time.Second))

	// DB Connect
	repo, err := mongo.NewRepository("mongodb://localhost:27017", "url-shortener", "admin", "password", ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Repository and logic layer
	logic := service.NewRedirectService(repo)
	api := handler.NewHandler(ctx, logic)

	// Router
	h := mux.NewRouter()
	h.HandleFunc("/api/v1/redirect", api.CreateRedirect()).Methods(http.MethodPost)
	h.HandleFunc("/api/v1/redirect/{code}", api.FindRedirectByCode()).Methods(http.MethodGet)

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
