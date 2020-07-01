package types

import (
	"github.com/adigunhammedolalekan/microservices-test/destroyer/proto/pb"
	"github.com/google/uuid"
	"time"
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
	Id        string    `json:"id"`
	Message   string    `json:"message"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

func NewEvent(name string, data []*Target) *Event {
	return &Event{Id: uuid.New().String(), Name: name, Data: data}
}

func NewTarget(t *pb.Target) (*Target, error) {
	c, err := time.Parse(time.RFC3339, t.CreatedOn)
	if err != nil {
		return nil, err
	}
	u, err := time.Parse(time.RFC3339, t.UpdatedOn)
	if err != nil {
		return nil, err
	}
	return &Target{
		Id:        t.Id,
		Message:   t.Message,
		CreatedOn: c,
		UpdatedOn: u,
	}, nil
}

func Convert(targets []*Target) []*pb.Target {
	data := make([]*pb.Target, 0, len(targets))
	for _, next := range targets {
		t := &pb.Target{
			Id: next.Id, Message: next.Message, CreatedOn: next.CreatedOn.String(), UpdatedOn: next.UpdatedOn.String(),
		}
		data = append(data, t)
	}
	return data
}
