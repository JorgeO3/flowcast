package mongo

import "time"

// Option is a function that sets a configuration option.
type Option func(*DB)

// MaxPoolSize is a function that sets a configuration option.
func MaxPoolSize(size uint64) Option {
	return func(c *DB) {
		c.maxPoolSize = size
	}
}

// ConnAttempts is a function that sets a configuration option.
func ConnAttempts(attempts int) Option {
	return func(c *DB) {
		c.connAttempts = attempts
	}
}

// ConnTimeout is a function that sets a configuration option.
func ConnTimeout(timeout time.Duration) Option {
	return func(c *DB) {
		c.connTimeout = timeout
	}
}
