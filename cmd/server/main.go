package main

import (
	"github.com/ahmetcancicek/go-url-shortener/repository/mongo"
	"github.com/ahmetcancicek/go-url-shortener/shortener"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	DBName   = "url-shortener"
	URI      = "mongodb://localhost:27017"
	USERNAME = "admin"
	PASSWORD = "password"
	TIMEOUT  = 10
)

func main() {

	//1- DB mongodb://localhost:27017
	repo, err := mongo.NewRepository(URI, USERNAME, PASSWORD, DBName, TIMEOUT)
	if err != nil {
		log.Fatal(err)
	}
	service := shortener.NewRedirectService(repo)
	handler := shortener.NewHandler(service)

	// 2- Router
	h := mux.NewRouter()
	h.HandleFunc("/v1/redirect", handler.CreateRedirect()).Methods(http.MethodPost)
	h.HandleFunc("/v1/redirect/{code}", handler.FindRedirectByCode()).Methods(http.MethodGet)

	// db.createUser({user:'admin',pwd:'password',roles:[{role:'readWrite',db:'url-shortener'}]})

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
