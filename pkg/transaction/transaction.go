// Package transaction provides abtractions for database transactions.
package transaction

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Manager is an interface that defines the methods to manage transactions.
type Manager interface {
	Begin(ctx context.Context) (Transaction, error)
}

// Transaction is an interface that defines the methods to manage a transaction.
type Transaction interface {
	Commit() error
	Rollback() error
}

// PgxTransactionManager is the implementation of the TransactionManager interface for the Pgx driver.
type PgxTransactionManager struct {
	Pool *pgxpool.Pool
}

// NewPgxTransactionManager creates a new instance of PgxTransactionManager.
func NewPgxTransactionManager(pool *pgxpool.Pool) *PgxTransactionManager {
	return &PgxTransactionManager{Pool: pool}
}

// Begin starts a new transaction.
func (tm *PgxTransactionManager) Begin(ctx context.Context) (Transaction, error) {
	tx, err := tm.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &PgxTransaction{Tx: tx}, nil
}

// PgxTransaction is the implementation of the Transaction interface for the Pgx driver.
type PgxTransaction struct {
	Tx pgx.Tx
}

// Commit commits the transaction.
func (t *PgxTransaction) Commit() error {
	return t.Tx.Commit(context.Background())
}

// Rollback rolls back the transaction.
func (t *PgxTransaction) Rollback() error {
	return t.Tx.Rollback(context.Background())
}
