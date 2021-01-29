package handler

import (
	"context"
	"encoding/json"
	"github.com/ahmetcancicek/go-url-shortener/internal/app/model"
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

		ctx := context.Background()

		redirect, err := h.redirectService.FindByCode(ctx, string(code))
		if err != nil {
			model.RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		// TODO: Should be updated click number
		model.RespondWithJSON(w, http.StatusOK, redirect)
	}
}

func (h handler) CreateRedirect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		redirect := &model.Redirect{}

		// 1. Decode request body
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&redirect); err != nil {
			model.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		// 2. Check data
		validate := validator.New()
		err := validate.Struct(redirect)
		if err != nil {
			model.RespondWithError(w, http.StatusBadRequest, err.Error())
		}

		ctx := context.Background()

		shortURL := shortid.MustGenerate()
		_, err = h.redirectService.FindByCode(ctx, shortURL)
		if err == nil {
			model.RespondWithError(w, http.StatusBadRequest, "Could not create a url!")
			return
		}

		// Generate ID
		redirect.Code = shortURL

		// 3. Save
		createdRedirect, err := h.redirectService.Save(ctx, redirect)

		model.RespondWithJSON(w, http.StatusOK, createdRedirect)
	}
}

func (h handler) CreateMultiRedirect() http.HandlerFunc {
	// TODO: We can do this operation for multiple url
	panic("implement me")
}
