package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
)

type sampleData struct {
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	projectID := os.Getenv("GCP_PROJECT_ID")
	topicName := os.Getenv("PUBSUB_TOPIC")
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create the client: %v", err)
	}

	data, err := json.Marshal(&sampleData{
		CreatedAt: time.Now().UTC(),
	})
	if err != nil {
		log.Fatalf("Invalid JSON: %v", err)
	}

	topic := client.Topic(topicName)
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(data),
	})

	msgID, err := result.Get(ctx)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	fmt.Printf("Message published: %v", msgID)
}
