// Package mongotx provides transaction management for MongoDB operations.
//
// It offers a flexible and robust way to execute functions within a MongoDB transaction context,
// handling session lifecycle and providing retry mechanisms for transient errors.
// This package is particularly useful for maintaining data consistency in distributed systems
// or when performing complex, multi-step database operations.
//
// Example usage:
//
//	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer client.Disconnect(ctx)
//
//	logger := logrus.New()
//	txManager, err := mongotx.New(client, mongotx.WithDefaultTxOptions(), mongotx.WithLogger(logger))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = txManager.Run(ctx, func(ctx mongo.SessionContext) error {
//		collection := client.Database("testdb").Collection("testcol")
//		_, err := collection.InsertOne(ctx, map[string]string{"name": "example"})
//		return err
//	})
//
//	if err != nil {
//		log.Fatalf("Transaction failed: %v", err)
//	} else {
//		log.Println("Transaction succeeded")
//	}
package mongotx

import (
	"context"
	"errors"
	"fmt"

	"github.com/JorgeO3/flowcast/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// ErrNilMongoClient is returned when a nil MongoDB client is provided.
var ErrNilMongoClient = errors.New("mongo client cannot be nil")

// TxManager defines the interface for transaction management operations.
// Implementations of this interface should provide methods to run operations
// within a transaction context.
type TxManager interface {
	Run(ctx context.Context, fn func(ctx context.Context) error) error
}

// MongoTx implements the TxManager interface for MongoDB.
// It encapsulates the MongoDB client, session options, transaction options,
// and a logger for comprehensive transaction management.
type MongoTx struct {
	client      *mongo.Client
	sessionOpts *options.SessionOptions
	txOpts      *options.TransactionOptions
	logger      logger.Interface
}

// TxOption is a function type used to apply optional configurations to MongoTx.
// This pattern allows for flexible and extensible configuration of the MongoTx struct.
type TxOption func(*MongoTx)

// WithDefaultTxOptions returns a TxOption that sets default transaction and session options.
// It uses majority read concern, primary read preference, and majority write concern.
func WithDefaultTxOptions() TxOption {
	return func(tx *MongoTx) {
		tx.txOpts = options.Transaction().
			SetReadConcern(readconcern.Majority()).
			SetWriteConcern(writeconcern.Majority()).
			SetReadPreference(readpref.Primary())
		tx.sessionOpts = options.Session().
			SetDefaultReadConcern(readconcern.Majority()).
			SetDefaultWriteConcern(writeconcern.Majority()).
			SetDefaultReadPreference(readpref.Primary())
	}
}

// WithTransactionOptions allows setting custom transaction options.
// This is useful when specific transaction behaviors are required.
func WithTransactionOptions(opts *options.TransactionOptions) TxOption {
	return func(tx *MongoTx) {
		tx.txOpts = opts
	}
}

// WithSessionOptions allows setting custom session options.
// This can be used to configure session-specific behaviors.
func WithSessionOptions(opts *options.SessionOptions) TxOption {
	return func(tx *MongoTx) {
		tx.sessionOpts = opts
	}
}

// WithLogger sets a custom logger for the MongoTx instance.
// This allows for integration with existing logging systems.
func WithLogger(logger logger.Interface) TxOption {
	return func(tx *MongoTx) {
		tx.logger = logger
	}
}

// New creates and returns a new TxManager (MongoTx) instance.
// It requires a valid MongoDB client and accepts optional TxOptions for configuration.
func New(client *mongo.Client, opts ...TxOption) (TxManager, error) {
	if client == nil {
		return nil, ErrNilMongoClient
	}
	tx := &MongoTx{client: client}
	for _, opt := range opts {
		opt(tx)
	}

	// Set default options if not provided
	if tx.txOpts == nil || tx.sessionOpts == nil {
		WithDefaultTxOptions()(tx)
	}

	return tx, nil
}

// Run executes the provided function within a MongoDB transaction.
// It handles the session lifecycle, transaction commit/abort logic,
// and recovers from panics within the function.
func (m *MongoTx) Run(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	session, err := m.client.StartSession(m.sessionOpts)
	if err != nil {
		m.logger.Error("Failed to start session: %v", err)
		return err
	}
	defer session.EndSession(ctx)

	defer func() {
		if r := recover(); r != nil {
			m.logger.Error("Panic occurred during transaction: %v. Aborting transaction.", r)
			_ = session.AbortTransaction(ctx)
			err = fmt.Errorf("panic occurred during transaction: %v", r)
		}
	}()

	_, err = session.WithTransaction(ctx, func(sc mongo.SessionContext) (interface{}, error) {
		return nil, fn(sc)
	}, m.txOpts)

	if err != nil {
		m.logger.Error("Transaction failed: %v", err)
	} else {
		m.logger.Info("Transaction completed successfully")
	}
	return err
}
