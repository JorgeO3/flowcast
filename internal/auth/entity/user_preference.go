package entity

// UserPref represents the user's notification preferences.
type UserPref struct {
	ID                 int
	UserID             int
	EmailNotifications bool
	SmsNotifications   bool
}

// NewUserPref creates a new instance of UserPref.
func NewUserPref(userID int, emailNotif, smsNotif bool) *UserPref {
	return &UserPref{
		UserID:             userID,
		EmailNotifications: emailNotif,
		SmsNotifications:   smsNotif,
	}
}
