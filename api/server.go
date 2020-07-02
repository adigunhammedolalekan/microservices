package api

import (
	"github.com/go-chi/chi"
	"net/http"
)

type ApiService struct {
	store  Store
	router *chi.Mux
}

func New(s Store) *ApiService {
	a := &ApiService{store: s}
	a.mountRoutes()
	return a
}

func (svc *ApiService) mountRoutes() {
	router := chi.NewRouter()
	handler := newApiHandler(svc.store)
	router.Post("/api/event/new", handler.CreateEventHandler)
	router.Get("/api/events", handler.ListEventsHandler)
	svc.router = router
}

func (svc *ApiService) Run(addr string) error {
	return http.ListenAndServe(addr, svc.router)
}
