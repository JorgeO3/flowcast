package redpanda

// BaseEvent represents a generic event structure for messaging.
type BaseEvent interface{}

// Admin defines the interface for administrative operations on Redpanda.
type Admin interface {
	// CreateTopic creates a new topic with the given name.
	// Returns an error if the operation fails.
	CreateTopic(topic string) error

	// TopicExists checks if a topic with the given name exists.
	// Returns true if the topic exists, false otherwise.
	// An error is returned if the check operation fails.
	TopicExists(topic string) (bool, error)

	// Close releases resources associated with the Admin client.
	// It's important to call this method when done using the Admin client.
	Close() error
}

// Producer defines the interface for publishing events to Redpanda.
type Producer interface {
	// Publish sends an event to the configured topic.
	// Returns an error if the publish operation fails.
	Publish(event BaseEvent, topic string) error

	// Close releases resources associated with the Producer.
	// It's important to call this method when done using the Producer.
	Close() error
}

// Consumer defines the interface for consuming events from Redpanda.
type Consumer interface {
	// Subscribe starts consuming events from the configured topic.
	// The provided handler function is called for each received event.
	// This method blocks until an error occurs or the consumer is closed.
	Subscribe(handler func(BaseEvent)) error

	// Close stops the consumer and releases associated resources.
	// It's important to call this method when done using the Consumer.
	Close() error
}
