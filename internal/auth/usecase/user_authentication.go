package usecase

import (
	"context"

	"github.com/asaskevich/govalidator"
	"gitlab.com/JorgeO3/flowcast/internal/auth/repository"
	"gitlab.com/JorgeO3/flowcast/pkg/logger"
)

// UserAuthenticationInput represents the input data required for user authentication.
type UserAuthenticationInput struct {
	Email    string
	Password string
}

// UserAuthenticationOutput represents the output data returned after user authentication.
type UserAuthenticationOutput struct {
	UserID    string
	AuthToken string
}

// UserAuthenticationUseCase handles the user authentication logic.
type UserAuthenticationUseCase struct {
	UserRepository repository.UserRepository
}

// NewUserAuthenticationUseCase creates a new instance of UserAuthenticationUseCase with the provided repository.
func NewUserAuthenticationUseCase(userRepository repository.UserRepository) *UserAuthenticationUseCase {
	return &UserAuthenticationUseCase{UserRepository: userRepository}
}

// Execute performs the user authentication and returns the result.
func (uc *UserAuthenticationUseCase) Execute(ctx context.Context, input UserAuthenticationInput, logg logger.Interface) (UserAuthenticationOutput, error) {

	// Validate input data
	ok, err := govalidator.ValidateStruct(input)
	if !ok {
		logg.Info("Invalid input data for user registration")
		return UserAuthenticationOutput{}, err
	}

	// Retrieve user details by username
	_, err = uc.UserRepository.FindByEmail(ctx, input.Email)
	if err != nil {
		logg.Error("failed to find user by username", "error", err)
		return UserAuthenticationOutput{}, err
	}

	return UserAuthenticationOutput{}, nil
}
