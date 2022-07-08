package main

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
)

func main() {
	if err := os.Setenv("PUBSUB_EMULATOR_HOST", "192.168.58.2:8085"); err != nil {
		log.Fatal(err)
	}
	client, err := pubsub.NewClient(context.Background(), "test")
	if err != nil {
		log.Fatalf("NewClient: %v", err)
	}
	for i := 0; i < 100; i++ {
		if _, err := client.CreateTopic(context.Background(), "cat123"); err != nil {
			log.Printf("CreateTopic: %v", err)
		}
		time.Sleep(time.Second)
	}
}
