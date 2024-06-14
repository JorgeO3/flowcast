package entity

import "time"

type (
	AuthStatus         string
	SubscriptionStatus string
	OauthProvider      string
	UserGender         string
)

const (
	AuthActive   AuthStatus = "active"
	AuthInactive AuthStatus = "inactive"
	AuthLocked   AuthStatus = "locked"

	SubscriptionActive    SubscriptionStatus = "active"
	SubscriptionInactive  SubscriptionStatus = "inactive"
	SubscriptionSuspended SubscriptionStatus = "suspended"

	GoogleProvider   OauthProvider = "google"
	GithubProvider   OauthProvider = "github"
	FacebookProvider OauthProvider = "facebook"

	MaleGender      UserGender = "male"
	FemaleGender    UserGender = "female"
	NonBinaryGender UserGender = "nonbinary"
	OtherGender     UserGender = "other"
)

type User struct {
	ID                 string
	Username           string
	Fullname           string
	Birthdate          time.Time
	Gender             UserGender
	Role               Role
	Email              string
	Phone              string
	Password           string
	AuthStatus         AuthStatus
	SubscriptionStatus SubscriptionStatus
	OAuthProvider      OauthProvider
	OAuthToken         string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func NewUser(id, username, password, email string) *User {
	return &User{
		ID:                 id,
		Username:           username,
		Password:           password,
		Email:              email,
		SubscriptionStatus: SubscriptionInactive,
		AuthStatus:         AuthInactive,
	}
}
