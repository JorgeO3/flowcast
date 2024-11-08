// Package kafka provides the Kafka consumer for the audsync service.
package kafka

import (
	"context"
	"encoding/json"

	"github.com/JorgeO3/flowcast/configs"
	"github.com/JorgeO3/flowcast/internal/audsync/events"
	apuc "github.com/JorgeO3/flowcast/internal/audsync/usecase/audprocess"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/redpanda"
)

// Consumer encapsulates the Redpanda consumer and related use cases.
type Consumer struct {
	consumer        redpanda.Consumer
	Logger          logger.Interface
	Cfg             *configs.AudsyncConfig
	CreateProcessUC *apuc.CreateProcessUC
	UpdateProcessUC *apuc.UpdateProcessUC
	DeleteProcessUC *apuc.DeleteProcessUC
}

// ConsumerOpts is a function type used for configuring the Consumer instance.
type ConsumerOpts func(*Consumer)

// WithConsumerLogger sets the logger for the consumer.
func WithConsumerLogger(logger logger.Interface) ConsumerOpts {
	return func(c *Consumer) {
		c.Logger = logger
	}
}

// WithConsumerConfig sets the configuration for the consumer.
func WithConsumerConfig(cfg *configs.AudsyncConfig) ConsumerOpts {
	return func(c *Consumer) {
		c.Cfg = cfg
	}
}

// WithConsumerCreateProcessUC sets the CreateProcess use case for the consumer.
func WithConsumerCreateProcessUC(uc *apuc.CreateProcessUC) ConsumerOpts {
	return func(c *Consumer) {
		c.CreateProcessUC = uc
	}
}

// WithConsumerUpdateProcessUC sets the UpdateProcess use case for the consumer.
func WithConsumerUpdateProcessUC(uc *apuc.UpdateProcessUC) ConsumerOpts {
	return func(c *Consumer) {
		c.UpdateProcessUC = uc
	}
}

// WithConsumerDeleteProcessUC sets the DeleteProcess use case for the consumer.
func WithConsumerDeleteProcessUC(uc *apuc.DeleteProcessUC) ConsumerOpts {
	return func(c *Consumer) {
		c.DeleteProcessUC = uc
	}
}

// NewConsumer creates a new Consumer instance with the provided brokers, topics, and configuration options.
func NewConsumer(brokers []string, topics []string, opts ...ConsumerOpts) (*Consumer, error) {
	consumer := &Consumer{}

	redpandaCons, err := redpanda.NewConsumer(brokers, topics)
	if err != nil {
		return nil, err
	}

	consumer.consumer = redpandaCons
	for _, opt := range opts {
		opt(consumer)
	}

	return consumer, nil
}

// handleCreateEvent processes CreateAudioProcessings events.
func (c *Consumer) handleCreateEvent(ctx context.Context, event events.BaseAudioEvent) error {
	var createEvent events.CreateAudioProcessingsEvent
	if err := json.Unmarshal(event.Payload, &createEvent); err != nil {
		c.Logger.Error("Error decoding create process event")
		return err
	}

	input := &apuc.CreateProcessInput{CreateAudioProcessingsEvent: createEvent}
	if _, err := c.CreateProcessUC.Execute(ctx, input); err != nil {
		c.Logger.Error("Error executing CreateProcess use case")
		return err
	}

	return nil
}

// handleUpdateEvent processes UpdateAudioProcessings events.
func (c *Consumer) handleUpdateEvent(ctx context.Context, event events.BaseAudioEvent) error {
	var updateEvent events.UpdateAudioProcessingsEvent
	if err := json.Unmarshal(event.Payload, &updateEvent); err != nil {
		c.Logger.Error("Error decoding update process event")
		return err
	}

	input := &apuc.UpdateProcessInput{UpdateAudioProcessingsEvent: updateEvent}
	if _, err := c.UpdateProcessUC.Execute(ctx, input); err != nil {
		c.Logger.Error("Error executing UpdateProcess use case")
		return err
	}

	return nil
}

// handleDeleteEvent processes DeleteAudioProcessings events.
func (c *Consumer) handleDeleteEvent(ctx context.Context, event events.BaseAudioEvent) error {
	var deleteEvent events.DeleteAudioProcessingsEvent
	if err := json.Unmarshal(event.Payload, &deleteEvent); err != nil {
		c.Logger.Error("Error decoding delete process event")
		return err
	}

	input := &apuc.DeleteProcessInput{DeleteAudioProcessingsEvent: deleteEvent}
	if _, err := c.DeleteProcessUC.Execute(ctx, input); err != nil {
		c.Logger.Error("Error executing DeleteProcess use case")
		return err
	}

	return nil
}

// Start begins message consumption and routes events to their appropriate handlers.
func (c *Consumer) Start(ctx context.Context) error {
	handler := func(event []byte) error {
		panic("implement me")
	}

	return c.consumer.Subscribe(handler)
}

// Close shuts down the consumer.
func (c *Consumer) Close() error {
	c.consumer.Close()
	return nil
}
