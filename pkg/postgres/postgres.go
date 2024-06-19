// Package postgres provides a series of utilities to interact with PostgreSQL through the Pgx driver.
package postgres

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Importa el paquete de
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

const (
	defaultMaxPoolSize  = 1
	defaultConnAttempts = 10
	defaultConnTimeout  = time.Second
)

// Postgres is a struct that holds the connection pool.
type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	Pool   *pgxpool.Pool
	config *pgxpool.Config
}

// New creates a new instance of Postgres.
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

func configurePool(connectionString string, pg *Postgres) (*pgxpool.Config, error) {
	poolConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = int32(pg.maxPoolSize)
	return poolConfig, nil
}

func connectWithRetry(pg *Postgres, poolConfig *pgxpool.Config) error {
	var connectionError error
	for pg.connAttempts > 0 {
		pg.Pool, connectionError = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if connectionError == nil {
			return nil
		}

		log.Printf("Postgres is trying to connect, attempts left: %d", pg.connAttempts-1)

		time.Sleep(pg.connTimeout)
		pg.connAttempts--
	}
	return fmt.Errorf("all connection attempts failed: %w", connectionError)
}

// Close closes the connection pool.
func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}

// RunMigrations runs the database migrations.
func (p *Postgres) RunMigrations(migrationsPath string, databaseName string) {
	db := stdlib.OpenDB(*p.Pool.Config().ConnConfig)
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not create driver: %v\n", err)
	}

	// Asegúrate de que migrationsPath tiene el esquema "file://"
	if !strings.HasPrefix(migrationsPath, "file://") {
		migrationsPath = "file://" + migrationsPath
	}

	// Inicialización de las migraciones
	m, err := migrate.NewWithDatabaseInstance(migrationsPath, databaseName, driver)
	if err != nil {
		log.Fatalf("Could not create migrate instance: %v\n", err)
	}

	// Ejecución de las migraciones
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not run up migrations: %v\n", err)
	}

	log.Println("Migrations ran successfully")
}
