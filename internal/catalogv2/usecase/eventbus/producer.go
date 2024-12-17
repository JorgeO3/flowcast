// Package eventbus defines the event producer contract.
package eventbus

import "context"

// Producer defines the methods that an event producer should implement.
type Producer interface {
	// Publish sends an event to the configured topic.
	// Returns an error if the publish operation fails.
	Publish(ctx context.Context, event any, topic string) error

	// Close releases resources associated with the Producer.
	// It's important to call this method when done using the Producer.
	Close() error
}
