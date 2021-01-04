package shortener

import "net/http"

type RedirectHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
}

type handler struct {
	redirectService RedirectService
}

func NewHandler(service RedirectService) RedirectHandler {
	return &handler{
		redirectService: service,
	}
}

func (h handler) Get(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func (h handler) Post(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func (h handler) Put(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}
