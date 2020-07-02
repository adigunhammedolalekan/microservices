package main

import (
	"fmt"
	"github.com/adigunhammedolalekan/microservices-test/api"
	"log"
	"os"
)

func main() {
	store, err := api.NewStore()
	if err != nil {
		log.Fatal(err)
	}
	svc := api.New(store)
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))

	log.Printf("api gateway running on %s", addr)
	if err := svc.Run(addr); err != nil {
		log.Fatal(err)
	}
}
