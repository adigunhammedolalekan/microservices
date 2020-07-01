package main

import (
	"fmt"
	"github.com/adigunhammedolalekan/microservices-test/destroyer"
	"github.com/adigunhammedolalekan/microservices-test/types"
	"github.com/apache/pulsar-client-go/pulsar"
	"net"
	"os"
)

func main() {
	svc, err := createDestroyerService()
	if err != nil {
		// ...
	}
	srv, err := destroyer.New(svc)
	if err != nil {
		// ...
	}
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		// ...
	}
	if err := srv.Run(lis); err != nil {

	}
}

func createDestroyerService() (*destroyer.Service, error) {
	pulsarClient, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: os.Getenv("PULSAR_URL"),
	})
	if err != nil {
		return nil, err
	}
	producer, err := pulsarClient.CreateProducer(pulsar.ProducerOptions{
		Topic: types.TopicName,
	})
	if err != nil {
		return nil, err
	}
	s, err := destroyer.NewStore()
	if err != nil {
		return nil, err
	}
	return destroyer.NewService(s, producer), nil
}
