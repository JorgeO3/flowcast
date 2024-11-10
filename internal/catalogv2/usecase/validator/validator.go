// Package validator defines the validator contract.
package validator

// Interface defines the methods that a validator must implement.
type Interface interface {
	// Validate validates the given input.
	// Returns an error if the input is invalid.
	Validate(input any) error
}
