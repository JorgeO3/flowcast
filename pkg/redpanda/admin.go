// Package redpanda provides a client library for interacting with Redpanda,
// a Kafka-compatible event streaming platform. It includes implementations
// for administrative operations, message production, and consumption.
package redpanda

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
)

// RedpAdmin is a Redpanda implementation of the Admin interface.
type RedpAdmin struct {
	client *kadm.Client
}

// NewAdmin creates an admin client that implements the Admin interface.
// It sets up a Kafka admin client using the provided broker addresses.
func NewAdmin(brokers []string) (Admin, error) {
	// Create a new Kafka client
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka client: %w", err)
	}

	// Create an admin client using the Kafka client
	admin := kadm.NewClient(client)
	return &RedpAdmin{client: admin}, nil
}

// TopicExists checks if a topic exists in the Kafka cluster.
func (a *RedpAdmin) TopicExists(topic string) (bool, error) {
	ctx := context.Background()
	topicsMetadata, err := a.client.ListTopics(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to list topics: %w", err)
	}

	for _, metadata := range topicsMetadata {
		if metadata.Topic == topic {
			return true, nil
		}
	}
	return false, nil
}

// CreateTopic creates a new topic in the Kafka cluster.
// It creates a topic with 1 partition and 1 replica.
func (a *RedpAdmin) CreateTopic(topic string) error {
	ctx := context.Background()
	resp, err := a.client.CreateTopics(ctx, 1, 1, nil, topic)
	if err != nil {
		return fmt.Errorf("failed to create topic: %w", err)
	}

	for _, ctr := range resp {
		if ctr.Err != nil {
			return fmt.Errorf("unable to create topic '%s': %w", ctr.Topic, ctr.Err)
		}
	}

	fmt.Printf("Created topic '%s'\n", topic)
	return nil
}

// Close gracefully shuts down the admin client and releases associated resources.
func (a *RedpAdmin) Close() error {
	a.client.Close()
	return nil
}
