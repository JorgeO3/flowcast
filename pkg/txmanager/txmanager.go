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

// CommandTag is an interface that abstracts pgconn.CommandTag.
type CommandTag interface {
	String() string
	RowsAffected() int64
}

// Rows is an interface that abstracts pgx.Rows.
type Rows interface {
	Next() bool
	Scan(dest ...interface{}) error
	Err() error
	Close()
}

// Row is an interface that abstracts pgx.Row.
type Row interface {
	Scan(dest ...interface{}) error
}

// Tx is an interface that defines the methods to manage a transaction.
type Tx interface {
	Begin(ctx context.Context) (Tx, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	Exec(ctx context.Context, query string, args ...interface{}) (CommandTag, error)
	Query(ctx context.Context, query string, args ...interface{}) (Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) Row
}

// PgxCommandTag wraps pgconn.CommandTag to implement CommandTag interface.
type PgxCommandTag struct {
	pgconn.CommandTag
}

// String returns the command tag as a string.
func (t PgxCommandTag) String() string {
	return t.CommandTag.String()
}

// RowsAffected returns the number of rows affected by the command.
func (t PgxCommandTag) RowsAffected() int64 {
	return t.CommandTag.RowsAffected()
}

// PgxRows wraps pgx.Rows to implement Rows interface.
type PgxRows struct {
	pgx.Rows
}

// PgxRow wraps pgx.Row to implement Row interface.
type PgxRow struct {
	pgx.Row
}

// PgxTx is the implementation of the Tx interface for the Pgx driver.
type PgxTx struct {
	Tx pgx.Tx
}

// Begin implements Tx.
func (t *PgxTx) Begin(ctx context.Context) (Tx, error) {
	tx, err := t.Tx.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &PgxTx{Tx: tx}, nil
}

// Commit commits the transaction.
func (t *PgxTx) Commit(ctx context.Context) error {
	return t.Tx.Commit(ctx)
}

// Rollback rolls back the transaction.
func (t *PgxTx) Rollback(ctx context.Context) error {
	return t.Tx.Rollback(ctx)
}

// Exec executes a query that doesn't return rows.
func (t *PgxTx) Exec(ctx context.Context, query string, args ...interface{}) (CommandTag, error) {
	tag, err := t.Tx.Exec(ctx, query, args...)
	return PgxCommandTag{CommandTag: tag}, err
}

// Query executes a query that returns rows.
func (t *PgxTx) Query(ctx context.Context, query string, args ...interface{}) (Rows, error) {
	rows, err := t.Tx.Query(ctx, query, args...)
	return &PgxRows{Rows: rows}, err
}

// QueryRow executes a query that returns a single row.
func (t *PgxTx) QueryRow(ctx context.Context, query string, args ...interface{}) Row {
	return &PgxRow{Row: t.Tx.QueryRow(ctx, query, args...)}
}

// DB defines common database operations.
type DB interface {
	Begin(ctx context.Context) (Tx, error)
	Exec(ctx context.Context, query string, args ...interface{}) (CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) Row
}

// TxManager defines transaction management operations.
type TxManager interface {
	Begin(ctx context.Context) (Tx, error)
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
func (m *PgxTxManager) Begin(ctx context.Context) (Tx, error) {
	tx, err := m.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &PgxTx{Tx: tx}, nil
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
