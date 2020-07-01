package types

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Data []*Target `json:"data"`
}

type Target struct {
	Id string `json:"id"`
	Message string `json:"message"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

func NewEvent(name string, data []*Target) *Event {
	return &Event{Id: uuid.New().String(), Name: name, Data: data}
}

func NewTarget(message string) *Target {
	return &Target{
		Id:        uuid.New().String(),
		Message:   message,
		CreatedOn: time.Now(),
		UpdatedOn: time.Now(),
	}
}
