package main

import (
	"github.com/adigunhammedolalekan/microservices-test/dealthstar"
	"github.com/adigunhammedolalekan/microservices-test/types"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	svc, err := createDeathStarService()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Println("dealthstar service running")
		if err := svc.Run(); err != nil {
			log.Fatal(err)
		}
	}()
	// quit on sigInt or sigTerm
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Println("received quit signal. Closing dealthstar...")
}

func createDeathStarService() (*dealthstar.Service, error) {
	pulsarClient, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: os.Getenv("PULSAR_URL"),
	})
	if err != nil {
		return nil, err
	}
	consumer, err := pulsarClient.Subscribe(pulsar.ConsumerOptions{
		Topic:            types.TopicName,
		SubscriptionName: types.TopicName,
		Type:             pulsar.Shared,
	})
	if err != nil {
		return nil, err
	}
	s, err := dealthstar.NewStore()
	if err != nil {
		return nil, err
	}
	return dealthstar.NewService(s, consumer), nil
}
