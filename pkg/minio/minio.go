// Package minio proporciona una serie de utilidades para interactuar con MinIO a través del SDK oficial.
package minio

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Constantes por defecto para la configuración del cliente.
const (
	defaultConnAttempts = 5
	defaultConnTimeout  = 2 * time.Second
)

// Client es una estructura que contiene la configuración y el cliente de MinIO.
type Client struct {
	endpoint      string
	accessKey     string
	secretKey     string
	useSSL        bool
	connAttempts  int
	connTimeout   time.Duration
	minioClient   *minio.Client
	clientOptions *minio.Options
}

// Option es un tipo de función que modifica la configuración del cliente.
type Option func(*Client)

// WithCredentials establece las credenciales de acceso.
func WithCredentials(accessKey, secretKey string) Option {
	return func(c *Client) {
		c.accessKey = accessKey
		c.secretKey = secretKey
	}
}

// WithSSL habilita o deshabilita el uso de SSL.
func WithSSL(useSSL bool) Option {
	return func(c *Client) {
		c.useSSL = useSSL
	}
}

// WithConnAttempts establece el número de intentos de conexión.
func WithConnAttempts(attempts int) Option {
	return func(c *Client) {
		c.connAttempts = attempts
	}
}

// WithConnTimeout establece el tiempo de espera entre intentos de conexión.
func WithConnTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.connTimeout = timeout
	}
}

// New crea una nueva instancia del cliente de MinIO.
func New(endpoint string, options ...Option) (*Client, error) {
	client := &Client{
		endpoint:     endpoint,
		useSSL:       true, // Valor por defecto
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}

	// Aplicar opciones proporcionadas
	for _, option := range options {
		option(client)
	}

	// Configurar las opciones del cliente
	client.clientOptions = &minio.Options{
		Creds:  credentials.NewStaticV4(client.accessKey, client.secretKey, ""),
		Secure: client.useSSL,
	}

	// Intentar conectar con reintentos
	err := client.connectWithRetry()
	if err != nil {
		return nil, fmt.Errorf("minio - New - connectWithRetry: %w", err)
	}

	return client, nil
}

// connectWithRetry intenta establecer una conexión con MinIO con reintentos.
func (c *Client) connectWithRetry() error {
	var connectionError error

	for c.connAttempts > 0 {
		minioClient, err := minio.New(c.endpoint, c.clientOptions)
		if err != nil {
			connectionError = err
			log.Printf("MinIO está intentando conectarse, intentos restantes: %d", c.connAttempts-1)
			time.Sleep(c.connTimeout)
			c.connAttempts--
			continue
		}

		// Verificar la conexión realizando una operación sencilla, como listar los buckets
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err = minioClient.ListBuckets(ctx)
		if err != nil {
			connectionError = err
			log.Printf("MinIO está intentando conectarse, intentos restantes: %d", c.connAttempts-1)
			time.Sleep(c.connTimeout)
			c.connAttempts--
			continue
		}

		c.minioClient = minioClient
		return nil
	}

	return fmt.Errorf("todos los intentos de conexión fallaron: %w", connectionError)
}

// Close desconecta el cliente de MinIO si es necesario.
func (c *Client) Close() {
	// El cliente de MinIO no requiere una desconexión explícita,
	// pero aquí puedes agregar lógica de limpieza si es necesario.
}

// GetClient devuelve el cliente de MinIO para realizar operaciones adicionales.
func (c *Client) GetClient() *minio.Client {
	return c.minioClient
}
