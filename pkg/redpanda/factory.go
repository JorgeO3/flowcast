package redpanda

import (
	"fmt"
)

// Factory is responsible for creating Admin, Producer, and Consumer instances
// based on a given configuration.
type Factory struct {
	config Config
}

// NewFactory creates a new Factory instance with the provided configuration.
func NewFactory(config Config) *Factory {
	return &Factory{config: config}
}

// CreateAdmin instantiates a new Admin client using the factory's configuration.
// It returns an Admin interface and any error encountered during creation.
func (f *Factory) CreateAdmin() (Admin, error) {
	admin, err := NewAdmin(f.config.Brokers)
	if err != nil {
		return nil, fmt.Errorf("failed to create Admin: %w", err)
	}
	return admin, nil
}

// CreateProducer instantiates a new Producer client for the specified topic
// using the factory's configuration.
// It returns a Producer interface and any error encountered during creation.
func (f *Factory) CreateProducer() (Producer, error) {
	producer, err := NewProducer(f.config.Brokers)
	if err != nil {
		return nil, fmt.Errorf("failed to create Producer: %w", err)
	}
	return producer, nil
}

// CreateConsumer instantiates a new Consumer client for the specified topic
// using the factory's configuration.
// It returns a Consumer interface and any error encountered during creation.
func (f *Factory) CreateConsumer(topic string) (Consumer, error) {
	consumer, err := NewConsumer(f.config.Brokers, topic)
	if err != nil {
		return nil, fmt.Errorf("failed to create Consumer: %w", err)
	}
	return consumer, nil
}
