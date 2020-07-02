package api

import (
	"bytes"
	"encoding/json"
	"github.com/adigunhammedolalekan/microservices-test/types"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mockTargets = []*types.Target{
		{Id: "1", Message: "hello"}, {Id: "2", Message: "world"},
	}
)
type mockStore struct {}

func (ms *mockStore) CreateEvent(e *types.Event) error {
	return nil
}
func (ms *mockStore) ListEvents() ([]*types.Target, error) {
	return mockTargets, nil
}


func TestApiHandler_CreateEventHandler(t *testing.T) {
	e := &types.Event{
		Id:   uuid.New().String(),
		Name: "targets.acquired",
		Data: mockTargets,
	}
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(e); err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/event/new", buf)

	s := &mockStore{}
	handler := newApiHandler(s)

	handler.CreateEventHandler(w, r)
	if code := w.Code; code != http.StatusOK {
		t.Fatalf("expected response code ok; got %d instead", code)
	}
}

func TestApiHandler_CreateEventHandlerBadRequest(t *testing.T) {
	buf := bytes.NewBufferString(`{invalid json data}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/event/new", buf)

	s := &mockStore{}
	handler := newApiHandler(s)

	handler.CreateEventHandler(w, r)
	if code := w.Code; code != http.StatusBadRequest {
		t.Fatalf("expected response code BadRequest; got %d instead", code)
	}
}

func TestApiHandler_ListEventsHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/events", nil)

	s := &mockStore{}
	handler := newApiHandler(s)

	handler.ListEventsHandler(w, r)
	if code := w.Code; code != http.StatusOK {
		t.Fatalf("expected response code ok; got %d instead", code)
	}
}
