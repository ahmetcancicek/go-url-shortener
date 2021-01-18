package handler

import (
	"encoding/json"
	"github.com/ahmetcancicek/go-url-shortener/internal/app/shortener"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/teris-io/shortid"
	"net/http"
)

type RedirectHandler interface {
	FindRedirectByCode() http.HandlerFunc
	CreateRedirect() http.HandlerFunc
	CreateMultiRedirect() http.HandlerFunc
}

type handler struct {
	redirectService shortener.RedirectService
}

func NewHandler(service shortener.RedirectService) RedirectHandler {
	return &handler{
		redirectService: service,
	}
}

func (h handler) FindRedirectByCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		code := vars["code"]

		redirect, err := h.redirectService.FindByCode(string(code))
		if err != nil {
			shortener.RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		// TODO: Should be updated click number
		shortener.RespondWithJSON(w, http.StatusOK, redirect)
	}
}

func (h handler) CreateRedirect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		redirect := &shortener.Redirect{}

		// 1. Decode request body
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&redirect); err != nil {
			shortener.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		// 2. Check data
		validate := validator.New()
		err := validate.Struct(redirect)
		if err != nil {
			shortener.RespondWithError(w, http.StatusBadRequest, err.Error())
		}

		shortURL := shortid.MustGenerate()
		_, err = h.redirectService.FindByCode(shortURL)
		if err == nil {
			shortener.RespondWithError(w, http.StatusBadRequest, "Could not create a url!")
			return
		}

		// Generate ID
		redirect.Code = shortURL

		// 3. Save
		createdRedirect, err := h.redirectService.Save(redirect)

		shortener.RespondWithJSON(w, http.StatusOK, createdRedirect)
	}
}

func (h handler) CreateMultiRedirect() http.HandlerFunc {
	// TODO: We can do this operation for multiple url
	panic("implement me")
}
