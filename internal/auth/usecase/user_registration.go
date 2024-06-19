package usecase

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"gitlab.com/JorgeO3/flowcast/internal/auth/entity"
	"gitlab.com/JorgeO3/flowcast/internal/auth/repository"
	"gitlab.com/JorgeO3/flowcast/pkg/logger"
	"gitlab.com/JorgeO3/flowcast/pkg/transaction"
)

// SocialLink represents a social media link in JSON
type SocialLink struct {
	Platform string `json:"platform" valid:"required,alphanum"`
	URL      string `json:"url" valid:"required,url"`
}

// UserRegInput represents the input data required for user registration.
type UserRegInput struct {
	Username    string       `json:"username" valid:"required,alphanum"`
	Email       string       `json:"email" valid:"required,email"`
	Password    string       `json:"password" valid:"required"`
	FullName    string       `json:"fullname" valid:"required"`
	Birthdate   time.Time    `json:"birthdate" valid:"required"`
	Gender      string       `json:"gender" valid:"required"`
	Phone       string       `json:"phone" valid:"required,numeric"`
	EmailNotif  bool         `json:"emailNotif" valid:"required"`
	SMSNotif    bool         `json:"smsNotif" valid:"required"`
	SocialLinks []SocialLink `json:"socialLinks"`
}

// UserRegOutput represents the output data returned after user registration.
type UserRegOutput struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// UserRegUC handles the user registration logic.
type UserRegUC struct {
	UserRepo  repository.UserRepo
	PrefRepo  repository.UserPrefRepo
	TxManager transaction.TxManager
}

// NewUserRegUC creates a new instance of UserRegUC.
func NewUserRegUC(userRepo repository.UserRepo, prefRepo repository.UserPrefRepo, txManager transaction.TxManager) *UserRegUC {
	return &UserRegUC{
		UserRepo:  userRepo,
		PrefRepo:  prefRepo,
		TxManager: txManager,
	}
}

// Execute performs the user registration.
func (uc *UserRegUC) Execute(ctx context.Context, input UserRegInput, logg logger.Interface) (UserRegOutput, error) {
	if err := validateInput(input); err != nil {
		logg.Info("Invalid input data for user registration", "error", err)
		return UserRegOutput{}, err
	}

	tx, err := uc.TxManager.Begin(ctx)
	if err != nil {
		logg.Error("Failed to start a new transaction", "error", err)
		return UserRegOutput{}, err
	}

	// Rollback the transaction if an error occurs in any of the following steps
	defer func() {
		if err != nil {
			tx.Rollback()
			logg.Error("Transaction rolled back", "error", err)
		}
	}()

	user, err := createUserEntity(input)
	if err != nil {
		logg.Error("Failed to create new user entity", "error", err)
		return UserRegOutput{}, err
	}

	userID, err := uc.UserRepo.Save(ctx, tx, user)
	if err != nil {
		logg.Error("Failed to save user to the database", "error", err)
		return UserRegOutput{}, err
	}

	userPreference := createUserPrefEntity(input, userID)
	if err := uc.PrefRepo.Save(ctx, tx, userPreference); err != nil {
		logg.Error("Failed to save user preference to the database", "error", err)
		return UserRegOutput{}, err
	}

	if err := tx.Commit(); err != nil {
		logg.Error("Failed to commit the transaction", "error", err)
		return UserRegOutput{}, err
	}

	return UserRegOutput{
		ID:       userID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

// validateInput validates the user registration input.
func validateInput(input UserRegInput) error {
	_, err := govalidator.ValidateStruct(input)
	return err
}

// createUserEntity creates a new user entity from the input data.
func createUserEntity(input UserRegInput) (*entity.User, error) {
	return entity.NewUser(
		input.Username,
		input.Email,
		input.Password,
		entity.WithFullName(input.FullName),
		entity.WithBirthdate(input.Birthdate),
		entity.WithGender(input.Gender),
		entity.WithPhone(input.Phone),
	)
}

// createUserPrefEntity creates a new user preference entity from the input data.
func createUserPrefEntity(input UserRegInput, userID int) *entity.UserPref {
	return entity.NewUserPref(userID, input.EmailNotif, input.SMSNotif)
}
