package hello

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/bigquery"
)

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// Item to be inserted on BigQuery.
type Item struct {
	ProcessedAt time.Time `bigquery:"inserted_at"`
	Data        string    `bigquery:"data"`
}

// PubSubBQ consumes a Pub/Sub message and inserts it into BigQuery.
func PubSubBQ(ctx context.Context, m PubSubMessage) error {
	data := string(m.Data)
	log.Println(data)

	client, err := bigquery.NewClient(ctx, "a2b-exp")
	if err != nil {
		log.Printf("Failed to create the client: %v", err)
		return err
	}
	defer client.Close()

	items := []*bigquery.StructSaver{}
	i := &bigquery.StructSaver{
		Struct: Item{
			ProcessedAt: time.Now().UTC(),
			Data:        data,
		},
	}
	items = append(items, i)

	uploader := client.Dataset("lucas_dataset_test").Table("table_test").Uploader()
	err = uploader.Put(ctx, items)
	if err != nil {
		log.Printf("Failed to insert items: %v", err)
		return err
	}

	log.Printf("Wrote to Big Query: %v", i)

	return nil
}
