// Package entity contains the definitions of the domain entities.
package entity

// SocialLink represents a social link entity.
type SocialLink struct {
	ID       int
	UserID   int
	Platform string
	URL      string
}

// NewSocialLink creates a new instance of SocialLink.
func NewSocialLink(userID int, platform, url string) *SocialLink {
	return &SocialLink{
		Platform: platform,
		URL:      url,
		UserID:   userID,
	}
}
