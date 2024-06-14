package entity

type UserPreference struct {
	ID                 uint
	UserID             uint
	EmailNotifications bool
	SmsNotifications   bool
}
