// Package txmanager provides transaction management for database operations.
package txmanager

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ErrNoTx is returned when no transaction is found in context.
var ErrNoTx = errors.New("no transaction in context")

// DB defines common database operations.
type DB interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

// TxManager defines transaction management operations.
type TxManager interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Run(ctx context.Context, fn func(context.Context) error) error
}

// PgxTxManager implements Manager for pgx.
type PgxTxManager struct {
	pool *pgxpool.Pool
}

// New creates a new PgxTxManager.
func New(pool *pgxpool.Pool) TxManager {
	return &PgxTxManager{pool: pool}
}

// Begin starts a new transaction.
func (m *PgxTxManager) Begin(ctx context.Context) (pgx.Tx, error) {
	tx, err := m.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Run executes fn within a transaction.
func (m *PgxTxManager) Run(ctx context.Context, fn func(context.Context) error) error {
	tx, err := m.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	err = fn(setTx(ctx, tx))
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

type txKey struct{}

// setTx adds tx to context.
func setTx(ctx context.Context, tx DB) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// GetTx retrieves tx from context.
func GetTx(ctx context.Context, fallback DB) DB {
	if tx, ok := ctx.Value(txKey{}).(DB); ok {
		return tx
	}
	return fallback
}
