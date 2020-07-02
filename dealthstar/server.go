package dealthstar

import (
	"context"
	"encoding/json"
	"github.com/adigunhammedolalekan/microservices-test/types"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
)

type Consumer interface {
	Receive(ctx context.Context) (pulsar.Message, error)
	Close()
}

type Service struct {
	store    Store
	consumer Consumer
}

func NewService(store Store, consumer Consumer) *Service {
	return &Service{store: store, consumer: consumer}
}

func (svc *Service) Run() error {
	defer func() {
		if err := svc.Close(); err != nil {
			log.Printf("service closing error: %s", err.Error())
		}
	}()
	log.Printf("consumer is running. ready to receive")
	for {
		message, err := svc.consumer.Receive(context.Background())
		if err != nil {
			continue
		}
		// move event creation to another goroutine to avoid
		// blocking the receive's goroutine
		go func(m pulsar.Message) {
			log.Printf("Processing new message. ID=%s", m.ID())
			p := m.Payload()
			event := &types.Event{}
			if err := json.Unmarshal(p, event); err != nil {
				log.Println("failed to bind event from json: ", err)
				return
			}
			if err := svc.store.CreateEvents(event); err != nil {
				log.Println("failed to persist event: ", err)
			}
		}(message)
	}
}

func (svc *Service) Close() error {
	log.Println("cleanup: closing destroyer service")
	svc.consumer.Close()
	return svc.store.Close()
}
