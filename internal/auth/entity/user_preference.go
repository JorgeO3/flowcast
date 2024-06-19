package entity

// UserPreference represents the user's notification preferences.
type UserPreference struct {
	ID                 uint
	UserID             uint
	EmailNotifications bool
	SmsNotifications   bool
}

// NewUserPreference creates a new instance of UserPreference.
func NewUserPreference(userID uint, emailNotif, smsNotif bool) *UserPreference {
	return &UserPreference{
		UserID:             userID,
		EmailNotifications: emailNotif,
		SmsNotifications:   smsNotif,
	}
}
