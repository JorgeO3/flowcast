package usecase

import (
	"context"

	"github.com/asaskevich/govalidator"
	"gitlab.com/JorgeO3/flowcast/configs"
	"gitlab.com/JorgeO3/flowcast/internal/auth/repository"
	"gitlab.com/JorgeO3/flowcast/pkg/logger"
)

// UserAuthInput represents the input data required for user authentication.
type UserAuthInput struct {
	Email    string
	Password string
}

// UserAuthOutput represents the output data returned after user authentication.
type UserAuthOutput struct {
	UserID    string
	AuthToken string
}

// UserAuthUC handles the user authentication logic.
type UserAuthUC struct {
	UserRepo repository.UserRepo
	Logg     logger.Interface
}

// NewUserAuthUC creates a new instance of UserAuthUC with the provided repository.
func NewUserAuthUC(userRepo repository.UserRepo, logg logger.Interface) *UserAuthUC {
	return &UserAuthUC{UserRepo: userRepo, Logg: logg}
}

// Execute performs the user authentication and returns the result.
func (uc *UserAuthUC) Execute(ctx context.Context, input UserAuthInput, cfg *configs.AuthConfig) (UserAuthOutput, error) {

	// Validate input data
	ok, err := govalidator.ValidateStruct(input)
	if !ok {
		uc.Logg.Info("Invalid input data for user authentication")
		return UserAuthOutput{}, err
	}

	// Retrieve user details by email
	_, err = uc.UserRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		uc.Logg.Error("Failed to find user by email", "error", err)
		return UserAuthOutput{}, err
	}

	return UserAuthOutput{}, nil
}
