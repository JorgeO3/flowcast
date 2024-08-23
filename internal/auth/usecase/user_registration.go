// Package usecase is responsible for handling the business logic of the authentication service.
package usecase

import (
	"context"
	"time"

	"github.com/JorgeO3/flowcast/configs"
	"github.com/JorgeO3/flowcast/internal/auth/entity"
	"github.com/JorgeO3/flowcast/internal/auth/errors"
	"github.com/JorgeO3/flowcast/internal/auth/repository"
	"github.com/JorgeO3/flowcast/internal/auth/service"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/txmanager"
	"github.com/asaskevich/govalidator"
)

// SocialLink represents a social media link in JSON
type SocialLink struct {
	Platform string `json:"platform" valid:"required,alphanum"`
	URL      string `json:"url" valid:"required,url"`
}

// UserRegistrationInput represents the input data required for user registration.
type UserRegistrationInput struct {
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

// UserRegistrationOutput represents the output data returned after user registration.
type UserRegistrationOutput struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// UserRegistrationUseCase handles the user registration logic.
type UserRegistrationUseCase struct {
	Mailer     service.Mailer
	UserRepo   repository.UserRepository
	PrefRepo   repository.UserPreferenceRepository
	SocialRepo repository.SocialLinkRepo
	EmailRepo  repository.EmailVerificationTokenRepo
	TxManager  txmanager.TxManager
	Logger     logger.Interface
}

// UserRegistrationUseCaseOption represents the options for the UserRegistrationUseCase.
type UserRegistrationUseCaseOption func(*UserRegistrationUseCase)

// NewUserRegistrationUseCase creates a new instance of UserRegistrationUseCase.
func NewUserRegistrationUseCase(opts ...UserRegistrationUseCaseOption) *UserRegistrationUseCase {
	uc := &UserRegistrationUseCase{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

// WithMailer adds a mailer to the UserRegistrationUseCase.
func WithMailer(mailer service.Mailer) UserRegistrationUseCaseOption {
	return func(uc *UserRegistrationUseCase) {
		uc.Mailer = mailer
	}
}

// WithUserRepo adds a user repository to the UserRegistrationUseCase.
func WithUserRepo(userRepo repository.UserRepository) UserRegistrationUseCaseOption {
	return func(uc *UserRegistrationUseCase) {
		uc.UserRepo = userRepo
	}
}

// WithUserPrefRepo adds a user preference repository to the UserRegistrationUseCase.
func WithUserPrefRepo(prefRepo repository.UserPreferenceRepository) UserRegistrationUseCaseOption {
	return func(uc *UserRegistrationUseCase) {
		uc.PrefRepo = prefRepo
	}
}

// WithSocialRepo adds a social link repository to the UserRegistrationUseCase.
func WithSocialRepo(socialRepo repository.SocialLinkRepo) UserRegistrationUseCaseOption {
	return func(uc *UserRegistrationUseCase) {
		uc.SocialRepo = socialRepo
	}
}

// WithEmailRepo adds an email repository to the UserRegistrationUseCase.
func WithEmailRepo(emailRepo repository.EmailVerificationTokenRepo) UserRegistrationUseCaseOption {
	return func(uc *UserRegistrationUseCase) {
		uc.EmailRepo = emailRepo
	}
}

// WithTxManager adds a transaction manager to the UserRegistrationUseCase.
func WithTxManager(txManager txmanager.TxManager) UserRegistrationUseCaseOption {
	return func(uc *UserRegistrationUseCase) {
		uc.TxManager = txManager
	}
}

// WithLogger adds a logger to the UserRegistrationUseCase.
func WithLogger(logg logger.Interface) UserRegistrationUseCaseOption {
	return func(uc *UserRegistrationUseCase) {
		uc.Logger = logg
	}
}

// Execute executes the user registration logic.
func (uc *UserRegistrationUseCase) Execute(ctx context.Context, input UserRegistrationInput, cfg *configs.AuthConfig) (*UserRegistrationOutput, error) {
	uc.Logger.Info("Starting user registration - input: %v", input)
	defer uc.Logger.Info("User registration process completed - username: %s", input.Username)

	if ok, err := govalidator.ValidateStruct(input); ok {
		uc.Logger.Warn("Invalid input data for user registration - error %s", err)
		return &UserRegistrationOutput{}, errors.NewValidation("invalid input data", err)
	}

	var user *entity.User

	// Start a transaction
	err := uc.TxManager.Run(ctx, func(ctx context.Context) error {
		var err error

		user, err = createUserEntity(input)
		if err != nil {
			uc.Logger.Error("Failed to create new user entity", "error", err)
			return errors.NewBadRequest("failed to create new user entity", err)
		}

		uc.Logger.Debug("Saving user to the database")

		user.ID, err = uc.UserRepo.Save(ctx, user)
		if err != nil {
			uc.Logger.Error("Failed to save user to the database - error: %s", err)
			return errors.NewInternal("failed to save user to the database", err)
		}

		uc.Logger.Debug("User saved successfully - userID: %s", user.ID)

		userPreference := createUserPrefEntity(input, user.ID)
		if err := uc.PrefRepo.Save(ctx, userPreference); err != nil {
			uc.Logger.Error("Failed to save user preference to the database - error: %s", err)
			return errors.NewInternal("failed to save user preference to the database", err)
		}

		uc.Logger.Debug("User preferences saved successfully - userID: %s", user.ID)

		socialLinks := createSocialLinkEntities(input, user.ID)
		uc.Logger.Debug("Saving social links to the database")
		if err := uc.SocialRepo.SaveTx(ctx, socialLinks); err != nil {
			return errors.NewInternal("failed to save social links to the database", err)
		}

		uc.Logger.Debug("Social links saved successfully - userID: %s - linkCount: %s", user.ID, len(socialLinks))

		return nil
	})

	if err != nil {
		uc.Logger.Error("Transaction failed - error: %s - user: %s", err, input.Username)
		return nil, err
	}

	// Check if user is nil after transaction avoids nil pointer dereference
	if user == nil {
		uc.Logger.Error("User is nil after transaction")
		return nil, errors.NewInternal("failed to create user", nil)
	}

	if err := uc.sendConfirmationEmail(cfg, user); err != nil {
		uc.Logger.Error("Failed to send confirmation email - error: %s", err)
		return nil, err
	}

	return &UserRegistrationOutput{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (uc *UserRegistrationUseCase) sendConfirmationEmail(cfg *configs.AuthConfig, user *entity.User) error {
	emailVerificationToken, err := createEmailVerificationToken(user.ID)
	if err != nil {
		uc.Logger.Error("Failed to create email verification token", "error", err)
		return errors.NewInternal("failed to create email verification token", err)
	}

	mailerConfig, err := createMailerConfig(cfg, user, emailVerificationToken)
	if err != nil {
		uc.Logger.Error("Failed to create mailer config", "error", err)
		return errors.NewInternal("failed to create mailer config", err)
	}

	if err := uc.Mailer.SendConfirmationEmail(mailerConfig); err != nil {
		uc.Logger.Error("Failed to send confirmation email", "error", err)
		return errors.NewInternal("failed to send confirmation email", err)
	}

	return nil
}
