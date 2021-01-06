package shortener

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type RedirectHandler interface {
	FindRedirectByCode() http.HandlerFunc
	CreateRedirect() http.HandlerFunc
}

type handler struct {
	redirectService RedirectService
}

func NewHandler(service RedirectService) RedirectHandler {
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
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, redirect)
	}
}

func (h handler) CreateRedirect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		redirect := &Redirect{}

		// 1. Decode request body
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&redirect); err != nil {
			RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		// 2. Check data
		validate := validator.New()
		err := validate.Struct(redirect)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
		}

		// Generate ID
		redirect.Code = "88454"
		redirect.CreatedAt = time.Now()
		redirect.Click = 0

		// 3. Save
		createdRedirect, err := h.redirectService.Save(redirect)

		RespondWithJSON(w, http.StatusOK, createdRedirect)
	}
}
