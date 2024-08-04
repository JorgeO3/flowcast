package usecase

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/asaskevich/govalidator"
	"gitlab.com/JorgeO3/flowcast/configs"
	"gitlab.com/JorgeO3/flowcast/internal/auth/entity"
	ae "gitlab.com/JorgeO3/flowcast/internal/auth/errors"
	"gitlab.com/JorgeO3/flowcast/internal/auth/repository"
	"gitlab.com/JorgeO3/flowcast/internal/auth/service"
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
	Mailer     service.Mailer
	UserRepo   repository.UserRepo
	PrefRepo   repository.UserPrefRepo
	SocialRepo repository.SocialLinkRepo
	EmailRepo  repository.EmailVerificationTokenRepo
	TxManager  transaction.TxManager
	Logger     logger.Interface
}

type UserRegUCOption func(*UserRegUC)

// NewUserRegUC creates a new instance of UserRegUC
func NewUserRegUC(options ...UserRegUCOption) *UserRegUC {
	uc := &UserRegUC{}
	for _, option := range options {
		option(uc)
	}
	return uc
}

// WithMailer -.
func WithMailer(mailer service.Mailer) UserRegUCOption {
	return func(uc *UserRegUC) {
		uc.Mailer = mailer
	}
}

// WithUserRepo -.
func WithUserRepo(userRepo repository.UserRepo) UserRegUCOption {
	return func(uc *UserRegUC) {
		uc.UserRepo = userRepo
	}
}

// WithUserPrefRepo -.
func WithUserPrefRepo(prefRepo repository.UserPrefRepo) UserRegUCOption {
	return func(uc *UserRegUC) {
		uc.PrefRepo = prefRepo
	}
}

// WithSocialRepo -.
func WithSocialRepo(socialRepo repository.SocialLinkRepo) UserRegUCOption {
	return func(uc *UserRegUC) {
		uc.SocialRepo = socialRepo
	}
}

// WithEmailRepo -.
func WithEmailRepo(emailRepo repository.EmailVerificationTokenRepo) UserRegUCOption {
	return func(uc *UserRegUC) {
		uc.EmailRepo = emailRepo
	}
}

// WithTxManager -.
func WithTxManager(txManager transaction.TxManager) UserRegUCOption {
	return func(uc *UserRegUC) {
		uc.TxManager = txManager
	}
}

// WithLogger -.
func WithLogger(logg logger.Interface) UserRegUCOption {
	return func(uc *UserRegUC) {
		uc.Logger = logg
	}
}

// Execute performs the user registration.
func (uc *UserRegUC) Execute(ctx context.Context, input UserRegInput, cfg *configs.AuthConfig) (UserRegOutput, ae.AuthError) {
	uc.Logger.Info("Starting user registration", "input", input)

	if err := validateInput(input); err != nil {
		uc.Logger.Warn("Invalid input data for user registration", "error", err)
		return UserRegOutput{}, ae.NewBadRequest("", err)
	}

	err := uc.TxManager.Transaction(ctx, func(ctx context.Context) error {
		userPreference := createUserPrefEntity(input, userID)
		if userPreference == nil {
			uc.Logger.Error("User preference entity is nil")
			return UserRegOutput{}, ae.NewBadRequest("")
		}

		if err := uc.PrefRepo.Save(ctx, tx, userPreference); err != nil {
			uc.Logger.Error("Failed to save user preference to the database", "error", err)
			return UserRegOutput{}, fmt.Errorf("usecase: failed to save user preference to the database: %w", err)
		}
	})

	// tx, err := uc.TxManager.Begin(ctx)
	// if err != nil {
	// 	uc.Logger.Error("Failed to start a new transaction", "error", err)
	// 	// return UserRegOutput{}, fmt.Errorf("usecase: failed to start a new transaction: %w", err)
	// 	return UserRegOutput{}, &DomainError{
	// 		Type: ErrorTypeInternal,
	// 	}
	// }

	// var commit = false

	// defer func() {
	// 	if !commit {
	// 		_ = tx.Rollback()
	// 		uc.Logger.Error("Transaction rolled back")
	// 	}
	// }()

	// user, err := createUserEntity(input)
	// if err != nil {
	// 	uc.Logger.Error("Failed to create new user entity", "error", err)
	// 	return UserRegOutput{}, fmt.Errorf("usecase: failed to create new user entity: %w", err)
	// }

	// uc.Logger.Debug("Saving user to the database")
	// userID, err := uc.UserRepo.Save(ctx, tx, user)
	// if err != nil {
	// 	uc.Logger.Error("Failed to save user to the database", "error", err)
	// 	return UserRegOutput{}, fmt.Errorf("usecase: failed to save user to the database: %w", err)
	// }

	// userPreference := createUserPrefEntity(input, userID)
	// if userPreference == nil {
	// 	uc.Logger.Error("User preference entity is nil")
	// 	return UserRegOutput{}, fmt.Errorf("usecase: user preference entity is nil")
	// }

	// uc.Logger.Debug("Saving user preference to the database")
	// if err := uc.PrefRepo.Save(ctx, tx, userPreference); err != nil {
	// 	uc.Logger.Error("Failed to save user preference to the database", "error", err)
	// 	return UserRegOutput{}, fmt.Errorf("usecase: failed to save user preference to the database: %w", err)
	// }

	// socialLinks := createSocialLinkEntities(input, userID)
	// uc.Logger.Debug("Saving social links to the database")
	// if err := uc.SocialRepo.SaveTx(ctx, tx, socialLinks); err != nil {
	// 	uc.Logger.Error("Failed to save social links to the database", "error", err)
	// 	return UserRegOutput{}, fmt.Errorf("usecase: failed to save social links to the database: %w", err)
	// }

	// emailVerificationToken, err := createEmailVerificationToken(userID)
	// if err != nil {
	// 	uc.Logger.Error("Failed to create email verification token", "error", err)
	// 	return UserRegOutput{}, fmt.Errorf("usecase: failed to create email verification token: %w", err)
	// }

	// uc.Logger.Debug("Saving email verification token to the database")
	// if err := uc.EmailRepo.SaveTx(ctx, tx, emailVerificationToken); err != nil {
	// 	uc.Logger.Error("Failed to save the email verification token to the database", "error", err)
	// 	return UserRegOutput{}, fmt.Errorf("usecase: failed to save the email verification token to the database: %w", err)
	// }

	// if err := tx.Commit(); err != nil {
	// 	uc.Logger.Error("Failed to commit the transaction", "error", err)
	// 	return UserRegOutput{}, fmt.Errorf("usecase: failed to commit the transaction: %w", err)
	// }
	// commit = true

	// uc.Logger.Info("User registration successful, sending confirmation email")
	// mailerConfig, err := createMailerConfig(cfg, user, emailVerificationToken)
	// if err != nil {
	// 	uc.Logger.Error("Failed to create mailer config", "error", err)
	// 	return UserRegOutput{}, fmt.Errorf("usecase: failed to create mailer config: %w", err)
	// }

	// if err := uc.Mailer.SendConfirmationEmail(mailerConfig); err != nil {
	// 	uc.Logger.Error("Failed to send confirmation email", "error", err)
	// 	return UserRegOutput{}, fmt.Errorf("usecase: failed to send confirmation email: %w", err)
	// }

	// return UserRegOutput{
	// 	ID:       userID,
	// 	Username: user.Username,
	// 	Email:    user.Email,
	// }, nil

	return UserRegOutput{
		ID:       1,
		Username: "",
		Email:    "",
	}, err
}

func validateInput(input UserRegInput) error {
	_, err := govalidator.ValidateStruct(input)
	return err
}

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

func createUserPrefEntity(input UserRegInput, userID int) *entity.UserPref {
	return entity.NewUserPref(userID, input.EmailNotif, input.SMSNotif)
}

func createSocialLinkEntities(input UserRegInput, userID int) []*entity.SocialLink {
	socialLinksLen := len(input.SocialLinks)
	socialLinks := make([]*entity.SocialLink, socialLinksLen)

	for i, socialLink := range input.SocialLinks {
		socialLinks[i] = entity.NewSocialLink(userID, socialLink.Platform, socialLink.URL)
	}
	return socialLinks
}

func createEmailVerificationToken(userID int) (*entity.EmailVerificationToken, error) {
	return entity.NewEmailVerificationToken(userID)
}

func createMailerConfig(cfg *configs.AuthConfig, user *entity.User, emailVer *entity.EmailVerificationToken) (*service.MailerConfig, error) {
	byteHTMLTemplate, err := os.ReadFile(cfg.EmailTemplate)
	if err != nil {
		return nil, err
	}

	htmlTemplate := string(byteHTMLTemplate)

	data := map[string]string{
		"UserName":        user.FullName,
		"ConfirmationURL": emailVer.Token,
	}

	return service.NewMailerConfig(data, user.Email, htmlTemplate, "email_template"), nil
}
