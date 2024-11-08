package kafka

import (
	"github.com/JorgeO3/flowcast/configs"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/redpanda"
)

// TODO: Fix the redpanda producer generic type

// Producer encapsulaes the redpanda consumer
type Producer struct {
	producer redpanda.Producer
	Logger   logger.Interface
	Cfg      *configs.AudsyncConfig
}

// ProducerOpts is used for configuring the Producer instance
type ProducerOpts func(*Producer)

// WithProducerLogger sets the logger for the producer
func WithProducerLogger(logger logger.Interface) ProducerOpts {
	return func(p *Producer) {
		p.Logger = logger
	}
}

// WithProducerConfig sets the configuration for the producer
func WithProducerConfig(cfg *configs.AudsyncConfig) ProducerOpts {
	return func(p *Producer) {
		p.Cfg = cfg
	}
}

// NewProducer creates a new Producer instance
func NewProducer(brokers []string, opts ...ProducerOpts) (*Producer, error) {
	// TODO: Fix the redpanda producer generic type
	rproducer, err := redpanda.NewProducer[string](brokers)
	if err != nil {
		return nil, err
	}

	producer := &Producer{producer: rproducer}

	for _, opt := range opts {
		opt(producer)
	}

	return producer, nil
}
