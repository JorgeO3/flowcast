package redpanda

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

// RedpProducer implements the Producer interface for Redpanda.
type RedpProducer struct {
	client *kgo.Client
}

// NewProducer creates a producer that implements the Producer interface.
// It sets up a Kafka client using the provided broker addresses and topic.
func NewProducer(brokers []string) (Producer, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka client: %w", err)
	}
	return &RedpProducer{client: client}, nil
}

// Publish sends an event to the specified topic.
// It serializes the event to JSON before sending.
func (p *RedpProducer) Publish(event BaseEvent, topic string) error {
	ctx := context.Background()

	// Serialize the event to JSON
	b, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	// Create a Kafka record
	record := &kgo.Record{
		Topic: topic,
		Value: b,
	}

	// Produce the record synchronously
	err = p.client.ProduceSync(ctx, record, nil).FirstErr()
	if err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}

	return nil
}

// Close gracefully shuts down the producer and releases associated resources.
func (p *RedpProducer) Close() error {
	p.client.Close()
	return nil
}
