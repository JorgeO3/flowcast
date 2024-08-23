// Package mongodb provides a series of utilities to interact with MongoDB through the official driver.
package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultMaxPoolSize  = 100
	defaultConnAttempts = 10
	defaultConnTimeout  = time.Second
)

// DB is a struct that holds the client connection.
type DB struct {
	maxPoolSize   uint64
	connAttempts  int
	connTimeout   time.Duration
	Client        *mongo.Client
	clientOptions *options.ClientOptions
}

// New creates a new instance of MongoDB.
func New(connectionString string, options ...Option) (*DB, error) {
	mdb := &DB{
		maxPoolSize:  defaultMaxPoolSize,
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}

	for _, option := range options {
		option(mdb)
	}

	clientOptions := configureClient(connectionString, mdb)
	err := connectWithRetry(mdb, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("mongodb - New - connectWithRetry: %w", err)
	}

	return mdb, nil
}

func configureClient(connectionString string, mdb *DB) *options.ClientOptions {
	clientOptions := options.Client().ApplyURI(connectionString)
	clientOptions.SetMaxPoolSize(mdb.maxPoolSize)
	return clientOptions
}

func connectWithRetry(mdb *DB, clientOptions *options.ClientOptions) error {
	var connectionError error
	for mdb.connAttempts > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		client, err := mongo.Connect(ctx, clientOptions)
		if err == nil {
			mdb.Client = client
			return nil
		}

		connectionError = err
		log.Printf("MongoDB is trying to connect, attempts left: %d", mdb.connAttempts-1)
		time.Sleep(mdb.connTimeout)
		mdb.connAttempts--
	}

	return fmt.Errorf("all connection attempts failed: %w", connectionError)
}

// Close disconnects the client.
func (m *DB) Close() {
	if m.Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		m.Client.Disconnect(ctx)
	}
}
