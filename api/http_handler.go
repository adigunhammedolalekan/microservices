package api

import (
	"encoding/json"
	"github.com/adigunhammedolalekan/microservices-test/api/errors"
	"github.com/adigunhammedolalekan/microservices-test/types"
	"github.com/go-chi/render"
	"net/http"
)

type apiHandler struct {
	store Store
}

func newApiHandler(store Store) *apiHandler {
	return &apiHandler{store: store}
}

func (handler *apiHandler) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	e := &types.Event{}
	if err := json.NewDecoder(r.Body).Decode(e); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.Respond(w, r, &Response{Error: true, Message: "bad request: malformed json body"})
		return
	}
	err := handler.store.CreateEvent(e)
	if err != nil {
		Respond(w, r, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.Respond(w, r, &Response{Error: false, Message: "event.stored"})
}

func (handler *apiHandler) ListEventsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := handler.store.ListEvents()
	if err != nil {
		Respond(w, r, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.Respond(w, r, &Response{Error: false, Message: "success", Data: data})
}

type Response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Respond(w http.ResponseWriter, r *http.Request, err error) {
	if e, ok := err.(*errors.ApiError); ok {
		render.Status(r, e.Code)
		render.Respond(w, r, &Response{Error: true, Message: e.Message})
	} else {
		render.Status(r, http.StatusInternalServerError)
		render.Respond(w, r, err.Error())
	}
}
