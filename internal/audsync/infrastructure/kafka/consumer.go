// Package kafka provides the kafka consumer for the audsync service.
package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/JorgeO3/flowcast/configs"
	apuc "github.com/JorgeO3/flowcast/internal/audsync/usecase/audprocess"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/redpanda"
)

// Consumer encapsula el consumidor de Redpanda y los casos de uso.
type Consumer struct {
	consumer redpanda.Consumer
	Logger   logger.Interface
	Cfg      *configs.AudsyncConfig

	GetProcessUC     *apuc.GetProcessUC
	CreateProcessUC  *apuc.CreateProcessUC
	UpdateProcessUC  *apuc.UpdateProcessUC
	DeleteProcessUC  *apuc.DeleteProcessUC
	GetManyProcessUC *apuc.GetManyProcessUC
}

// ConsumerOpts es una función que configura el consumidor.
type ConsumerOpts func(*Consumer)

// WithLogger configura el logger del consumidor.
func WithLogger(logger logger.Interface) ConsumerOpts {
	return func(c *Consumer) {
		c.Logger = logger
	}
}

// WithConfig configura la configuración del consumidor.
func WithConfig(cfg *configs.AudsyncConfig) ConsumerOpts {
	return func(c *Consumer) {
		c.Cfg = cfg
	}
}

// WithGetProcessUC configura el caso de uso GetProcess del consumidor.
func WithGetProcessUC(uc *apuc.GetProcessUC) ConsumerOpts {
	return func(c *Consumer) {
		c.GetProcessUC = uc
	}
}

// WithCreateProcessUC configura el caso de uso CreateProcess del consumidor.
func WithCreateProcessUC(uc *apuc.CreateProcessUC) ConsumerOpts {
	return func(c *Consumer) {
		c.CreateProcessUC = uc
	}
}

// WithUpdateProcessUC configura el caso de uso UpdateProcess del consumidor.
func WithUpdateProcessUC(uc *apuc.UpdateProcessUC) ConsumerOpts {
	return func(c *Consumer) {
		c.UpdateProcessUC = uc
	}
}

// WithDeleteProcessUC configura el caso de uso DeleteProcess del consumidor.
func WithDeleteProcessUC(uc *apuc.DeleteProcessUC) ConsumerOpts {
	return func(c *Consumer) {
		c.DeleteProcessUC = uc
	}
}

// WithGetManyProcessUC configura el caso de uso GetManyProcess del consumidor.
func WithGetManyProcessUC(uc *apuc.GetManyProcessUC) ConsumerOpts {
	return func(c *Consumer) {
		c.GetManyProcessUC = uc
	}
}

// NewConsumer crea una nueva instancia de KafkaConsumer.
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

// Start inicia el consumo de mensajes y delega el procesamiento a los casos de uso.
func (kc *Consumer) Start(ctx context.Context) error {
	handler := func(event redpanda.BaseEvent) {
		switch event.Type {
		case "create":
			var input apuc.CreateProcessInput

			if err := json.Unmarshal(event.Data, &input); err != nil {
				kc.Logger.Error("Error decoding create process event - err: ", err)
				return
			}

			if _, err := kc.CreateProcessUC.Execute(ctx, &input); err != nil {
				kc.Logger.Error("Error executing CreateProcess use case - err: ", err)
			}

		case "update":
		case "delete":
		default:
			log.Printf("Tipo de evento desconocido: %s", event.Type)
		}
	}

	return kc.consumer.Subscribe(handler)
}
