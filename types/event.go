package types

import (
	"github.com/adigunhammedolalekan/microservices-test/destroyer/proto/pb"
	"github.com/google/uuid"
)

const (
	TopicName = "targets.acquired"
)

type Event struct {
	Id   string    `json:"id"`
	Name string    `json:"name"`
	Data []*Target `json:"data"`
}

type Target struct {
	Id        string `json:"id"`
	Message   string `json:"message"`
	CreatedOn string `json:"created_on"`
	UpdatedOn string `json:"updated_on"`
}

func NewEvent(name string, data []*Target) *Event {
	return &Event{Id: uuid.New().String(), Name: name, Data: data}
}

func NewTarget(t *pb.Target) (*Target, error) {
	return &Target{
		Id:        t.Id,
		Message:   t.Message,
		CreatedOn: t.CreatedOn,
		UpdatedOn: t.UpdatedOn,
	}, nil
}

func Convert(targets []*Target) []*pb.Target {
	data := make([]*pb.Target, 0, len(targets))
	for _, next := range targets {
		t := &pb.Target{
			Id: next.Id, Message: next.Message, CreatedOn: next.CreatedOn, UpdatedOn: next.UpdatedOn,
		}
		data = append(data, t)
	}
	return data
}
