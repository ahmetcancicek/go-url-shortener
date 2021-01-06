package main

import (
	"github.com/ahmetcancicek/go-url-shortener/repository/mongo"
	"github.com/ahmetcancicek/go-url-shortener/shortener"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

const (
	DBName = "url-shortener"
	URI    = "mongodb://127.19.0.2:27017"
)

func main() {

	credential := options.Credential{
		Username: "admin",
		Password: "password",
	}

	//1- DB mongodb://localhost:27017
	repo, err := mongo.NewRepository(URI, DBName, credential, 10)
	if err != nil {
		log.Fatal(err)
	}
	service := shortener.NewRedirectService(repo)
	handler := shortener.NewHandler(service)

	// 2- Router
	h := mux.NewRouter()
	h.HandleFunc("/v1/redirect", handler.CreateRedirect()).Methods(http.MethodPost)

	// 3- Server
	srv := &http.Server{
		Addr:    ":8500",
		Handler: h,
	}

	// 4- Listen
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
