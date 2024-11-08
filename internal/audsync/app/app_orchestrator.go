package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/JorgeO3/flowcast/internal/audsync/infrastructure/kafka"
	"github.com/JorgeO3/flowcast/pkg/logger"
)

// Service represents a service that can be started and stopped
type Service interface {
	Start(context.Context) error
	Stop(context.Context) error
}

// HTTPServer encapsulate the raw HTTP server
type HTTPServer struct {
	server *http.Server
	logger logger.Interface
}

// NewHTTPServer is a constructuror for HTTPServer
func NewHTTPServer(addr string, handler http.Handler, logger logger.Interface) *HTTPServer {
	return &HTTPServer{
		server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
		logger: logger,
	}
}

// Start starts the HTTP server
func (s *HTTPServer) Start(_ context.Context) error {
	s.logger.Info("Starting HTTP server on " + s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Stop stops the HTTP server
func (s *HTTPServer) Stop(ctx context.Context) error {
	s.logger.Info("Shutting down HTTP server")
	return s.server.Shutdown(ctx)
}

// KafkaConsumerService encapsulates the Kafka consumer
type KafkaConsumerService struct {
	consumer *kafka.Consumer
	logger   logger.Interface
}

// NewKafkaConsumerService is a constructor for KafkaConsumerService
func NewKafkaConsumerService(consumer *kafka.Consumer, logger logger.Interface) *KafkaConsumerService {
	return &KafkaConsumerService{
		consumer: consumer,
		logger:   logger,
	}
}

// Start starts the Kafka consumer
func (s *KafkaConsumerService) Start(ctx context.Context) error {
	s.logger.Info("Starting Kafka consumer")
	return s.consumer.Start(ctx)
}

// Stop stops the Kafka consumer
func (s *KafkaConsumerService) Stop(_ context.Context) error {
	s.logger.Info("Stopping Kafka consumer")
	return s.consumer.Close()
}

// Orchestrator handles the lifecycle of multiple services
type Orchestrator struct {
	services []Service
	timeout  time.Duration
	logger   logger.Interface
	wg       sync.WaitGroup
	ctx      context.Context
	cancel   context.CancelFunc
}

// NewOrchestrator is a constructor for Orchestrator
func NewOrchestrator(timeout time.Duration, logger logger.Interface) *Orchestrator {
	ctx, cancel := context.WithCancel(context.Background())
	return &Orchestrator{
		timeout: timeout,
		logger:  logger,
		ctx:     ctx,
		cancel:  cancel,
	}
}

// AddService adds a service to the orchestrator
func (o *Orchestrator) AddService(service Service) {
	o.services = append(o.services, service)
}

// Start starts all services and waits for a shutdown signal
func (o *Orchestrator) Start() error {
	// Canal para señales del sistema operativo
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Iniciar todos los servicios
	for _, svc := range o.services {
		o.wg.Add(1)
		go func(s Service) {
			defer o.wg.Done()
			if err := s.Start(o.ctx); err != nil {
				o.logger.Error("Service error", "error", err)
				o.cancel() // Cancela todos los servicios si uno falla
			}
		}(svc)
	}

	// Esperar señal de shutdown
	<-sigChan
	o.logger.Info("Received shutdown signal")
	return o.shutdown()
}

func (o *Orchestrator) shutdown() error {
	// Crear contexto con timeout para el shutdown
	ctx, cancel := context.WithTimeout(context.Background(), o.timeout)
	defer cancel()

	// Cancelar el contexto principal
	o.cancel()

	// Detener todos los servicios
	for _, svc := range o.services {
		if err := svc.Stop(ctx); err != nil {
			o.logger.Error("Error stopping service", "error", err)
		}
	}

	// Esperar a que todos los servicios se detengan
	done := make(chan struct{})
	go func() {
		o.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		o.logger.Info("All services stopped successfully")
		return nil
	case <-ctx.Done():
		return errors.New("shutdown timeout exceeded")
	}
}
