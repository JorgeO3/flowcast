package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	pgx "github.com/jackc/pgx/v5/pgxpool"
)

const (
	defaultMaxPoolSize  = 1
	defaultConnAttempts = 10
	defaultConnTimeout  = time.Second
)

type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	pool *pgx.Pool
}

func New(connectionString string, options ...Option) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:  defaultMaxPoolSize,
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}

	for _, option := range options {
		option(pg)
	}

	poolConfig, err := configurePool(connectionString, pg)
	if err != nil {
		return nil, fmt.Errorf("postgres - New - configurePool: %w", err)
	}

	err = connectWithRetry(pg, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("postgres - New - connectWithRetry: %w", err)
	}

	return pg, nil
}

func configurePool(connectionString string, pg *Postgres) (*pgx.Config, error) {
	poolConfig, err := pgx.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = int32(pg.maxPoolSize)
	return poolConfig, nil
}

func connectWithRetry(pg *Postgres, poolConfig *pgx.Config) error {
	var connectionError error
	for pg.connAttempts > 0 {
		pg.pool, connectionError = pgx.NewWithConfig(context.Background(), poolConfig)
		if connectionError == nil {
			return nil
		}

		log.Printf("Postgres is trying to connect, attempts left: %d", pg.connAttempts-1)

		time.Sleep(pg.connTimeout)
		pg.connAttempts--
	}
	return fmt.Errorf("all connection attempts failed: %w", connectionError)
}

func (p *Postgres) Close() {
	if p.pool != nil {
		p.pool.Close()
	}
}
