package api

import (
	"context"
	"github.com/adigunhammedolalekan/microservices-test/api/errors"
	"github.com/adigunhammedolalekan/microservices-test/destroyer/proto/pb"
	"github.com/adigunhammedolalekan/microservices-test/types"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

type Store interface {
	CreateEvent(event *types.Event) error
	ListEvents() ([]*types.Target, error)
}

type destroyerStore struct {
	client pb.DestroyerServiceClient
}

func (s *destroyerStore) ListEvents() ([]*types.Target, error) {
	r, err := s.client.ListTargets(context.Background(), &pb.ListTargetsRequest{})
	if err != nil {
		log.Printf("listTargets error=%s", err.Error())
		return nil, errors.New(http.StatusInternalServerError, "failed to list targets. "+err.Error())
	}
	data := make([]*types.Target, 0)
	for _, next := range r.Data {
		t, err := types.NewTarget(next)
		if err == nil {
			data = append(data, t)
		}
	}
	return data, nil
}

func (s *destroyerStore) CreateEvent(event *types.Event) error {
	req := &pb.EventRequest{
		Id:   uuid.New().String(),
		Name: "targets.acquired",
		Data: types.Convert(event.Data),
	}
	r, err := s.client.AcquireTargets(context.Background(), req)
	if err != nil {
		log.Printf("acquireTargets error=%s", err.Error())
		return errors.New(http.StatusInternalServerError, "failed to create event at this time. Please retry later")
	}
	log.Println("event stored: ", r.MessageId)
	return nil
}

func NewStore() (Store, error) {
	conn, err := grpc.Dial(os.Getenv("DESTROYER_URL"), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := pb.NewDestroyerServiceClient(conn)
	return &destroyerStore{client: client}, nil
}
