// Package eventproducer defines the event producer contract.
package eventproducer

import "context"

// Interface defines the methods that an event producer should implement.
type Interface interface {
	// Publish sends an event to the configured topic.
	// Returns an error if the publish operation fails.
	Publish(ctx context.Context, event any, topic string) error

	// Close releases resources associated with the Producer.
	// It's important to call this method when done using the Producer.
	Close() error
}
