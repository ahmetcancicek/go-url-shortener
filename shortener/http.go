package shortener

import "net/http"

type RedirectHandler interface {
	FindRedirectByID() http.HandlerFunc
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

func (h handler) FindRedirectByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h handler) CreateRedirect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
