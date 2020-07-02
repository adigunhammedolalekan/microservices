package main

import (
	"fmt"
	"github.com/adigunhammedolalekan/microservices-test/destroyer"
	"github.com/adigunhammedolalekan/microservices-test/types"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
	"net"
	"os"
)

func main() {
	svc, err := createDestroyerService()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := svc.Close(); err != nil {
			log.Printf("failed to close destroyer service: %s", err.Error())
		}
	}()

	srv, err := destroyer.New(svc)
	if err != nil {
		log.Fatal(err)
	}
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("destroyer service running on %s", addr)
	if err := srv.Run(lis); err != nil {
		log.Fatal(err)
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
