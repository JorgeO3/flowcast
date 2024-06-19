// Package entity contains the definitions of the domain entities.
package entity

// SocialLink represents a social link entity.
type SocialLink struct {
	ID       int
	UserID   int
	Platform string
	URL      string
}
