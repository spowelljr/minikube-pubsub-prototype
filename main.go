package main

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"google.golang.org/grpc"
)

func main() {
	addr := "192.168.58.2:8085"
	os.Setenv("PUBSUB_EMULATOR_HOST", addr)
	_, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial: %v", err)
	}
	log.Println("dial successful")
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		client, err := pubsub.NewClient(context.Background(), "test")
		if err != nil {
			log.Printf("NewClient: %v", err)
			continue
		}
		if _, err := client.CreateTopic(context.Background(), "cat123"); err != nil {
			log.Printf("CreateTopic: %v", err)
		}
	}
}
