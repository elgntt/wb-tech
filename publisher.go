package main

import (
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("model.json")
	if err != nil {
		log.Fatal(err)
	}

	clusterID := "test-cluster"
	clientID := "test-client2"
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatal(err)
	}

	err = sc.Publish("orders", data) // does not return until an ack has been received from NATS Streaming
	if err != nil {
		log.Fatal(err)
	}

	err = sc.Close()
	if err != nil {
		log.Fatal(err)
	}
}
