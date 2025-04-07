package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zillow/zkafka"
)

// MyProcessor is our custom Kafka message processor
type MyProcessor struct{}

// Process processes each Kafka message
func (p *MyProcessor) Process(ctx context.Context, msg *zkafka.Message) error {
	log.Printf("Received message: topic=%s partition=%d offset=%d key=%s value=%s",
		msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	return nil
}

func main() {
	// Configure the Kafka consumer with zkafka
	cfg := zkafka.Config{
		Brokers:    []string{"broker:29092"},  // Kafka broker from Docker Compose
		GroupID:    "example-group",           // Consumer group ID
		Topics:     []string{"example-topic"}, // Kafka topic to consume
		NumWorkers: 4,                         // Number of worker goroutines for parallel processing
		Logger:     log.Default(),
	}

	// Create a new zkafka consumer instance
	processor := &MyProcessor{}
	zk, err := zkafka.New(context.Background(), cfg, processor)
	if err != nil {
		log.Fatalf("Error creating zkafka consumer: %v", err)
	}

	fmt.Println("Starting Kafka consumer...")

	// Run the consumer to start processing messages
	if err := zk.Run(context.Background()); err != nil {
		log.Fatalf("Error running zkafka: %v", err)
	}
}
