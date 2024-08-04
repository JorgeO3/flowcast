// Package transaction provides abstractions for database transactions.
package transaction

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/JorgeO3/flowcast/internal/auth/errors"
)

// TxManager is an interface that defines the methods to manage transactions.
type TxManager interface {
	Begin(ctx context.Context) (Tx, error)
	Transaction(ctx context.Context, f func(context.Context) errors.AuthError) errors.AuthError
}

// Tx is an interface that defines the methods to manage a transaction.
type Tx interface {
	Commit() error
	Rollback() error
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
}

// PgxTxManager is the implementation of the TxManager interface for the Pgx driver.
type PgxTxManager struct {
	Pool *pgxpool.Pool
}

// NewPgxTxManager creates a new instance of PgxTxManager.
func NewPgxTxManager(pool *pgxpool.Pool) *PgxTxManager {
	return &PgxTxManager{Pool: pool}
}

// Begin starts a new transaction.
func (p *PgxTxManager) Begin(ctx context.Context) (pgx.Tx, error) {
	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return &PgxTx{Tx: tx}, nil
}

type txKey struct{}

// injectTx injects transaction to context
func InjectTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// extractTx extracts transaction from context
func ExtractTx(ctx context.Context, pool *pgxpool.Pool) pgx.Tx {
	if tx, ok := ctx.Value(txKey{}).(Tx); ok {
		return tx.
	}
	return pool
}

func (p *PgxTxManager) Transaction(ctx context.Context, f func(context.Context) error) error {
	tx, err := p.Begin(ctx)
	if err != nil {
		return err
	}

	err = f(InjectTx(ctx, tx))
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// PgxTx is the implementation of the Tx interface for the Pgx driver.
type PgxTx struct {
	Tx pgx.Tx
}

// Commit commits the transaction.
func (t *PgxTx) Commit() error {
	return t.Tx.Commit(context.Background())
}

// Rollback rolls back the transaction.
func (t *PgxTx) Rollback() error {
	return t.Tx.Rollback(context.Background())
}

// Exec executes a query that doesn't return rows.
func (t *PgxTx) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return t.Tx.Exec(ctx, query, args...)
}

// Query executes a query that returns rows.
func (t *PgxTx) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return t.Tx.Query(ctx, query, args...)
}

// QueryRow executes a query that returns a single row.
func (t *PgxTx) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return t.Tx.QueryRow(ctx, query, args...)
}
