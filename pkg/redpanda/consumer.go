package redpanda

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/twmb/franz-go/pkg/kgo"
)

// ConsumerRedp is a Redpanda implementation of the Consumer interface.
type ConsumerRedp struct {
	client *kgo.Client
}

// NewConsumer creates a consumer that implements the Consumer interface.
// It sets up a Kafka client with the provided brokers and topic.
func NewConsumer(brokers []string, topics []string) (Consumer, error) {
	// Generate a unique group ID for this consumer instance
	groupID := uuid.New().String()

	// Create a new Kafka client with the specified configuration
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumerGroup(groupID),
		kgo.ConsumeTopics(topics...),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka client: %w", err)
	}

	return &ConsumerRedp{client: client}, nil
}

// Subscribe listens for messages in the topic and processes them using the provided handler.
func (c *ConsumerRedp) Subscribe(handler func([]byte) error) error {
	ctx := context.Background()

	for {
		fetches := c.client.PollFetches(ctx)
		if fetches.IsClientClosed() {
			return nil
		}

		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()

			// * NOTE: The deserialized message is handled by the provided handler function.
			if err := handler(record.Value); err != nil {
				fmt.Printf("Error processing message: %v\n", err)
				continue
			}
		}
	}
}

// Close gracefully shuts down the consumer and releases associated resources.
func (c *ConsumerRedp) Close() {
	c.client.Close()
}
