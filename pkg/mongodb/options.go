package mongodb

import "time"

// Option is a function that sets a configuration option.
type Option func(*DB)

// WithMaxPoolSize sets the maximum pool size.
func WithMaxPoolSize(size uint64) Option {
	return func(db *DB) {
		db.maxPoolSize = size
	}
}

// WithConnAttempts sets the number of connection attempts.
func WithConnAttempts(attempts int) Option {
	return func(db *DB) {
		db.connAttempts = attempts
	}
}

// WithConnTimeout sets the connection timeout.
func WithConnTimeout(timeout time.Duration) Option {
	return func(db *DB) {
		db.connTimeout = timeout
	}
}
