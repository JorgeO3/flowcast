package usecase

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"gitlab.com/JorgeO3/flowcast/internal/auth/entity"
	"gitlab.com/JorgeO3/flowcast/internal/auth/repository"
	"gitlab.com/JorgeO3/flowcast/pkg/logger"
)

// UserRegistrationInput represents the input data required for user registration.
type UserRegistrationInput struct {
	Username  string    `json:"username" valid:"required,alpha"`
	Email     string    `json:"email" valid:"required,email"`
	Password  string    `json:"password" valid:"required"`
	FullName  string    `json:"fullname" valid:"required"`
	Birthdate time.Time `json:"birthdate" valid:"required"`
	Gender    string    `json:"gender" valid:"required"`
	Phone     string    `json:"phone" valid:"required,numeric"`
}

// UserRegistrationOutput represents the output data returned after user registration.
type UserRegistrationOutput struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// UserRegistrationUseCase handles the user registration logic.
type UserRegistrationUseCase struct {
	UserRepository repository.UserRepository
}

// NewUserRegistrationUseCase creates a new instance of URUseCase.
func NewUserRegistrationUseCase(userRepository repository.UserRepository) *UserRegistrationUseCase {
	return &UserRegistrationUseCase{userRepository}
}

// Execute performs the user registration.
func (uc *UserRegistrationUseCase) Execute(ctx context.Context, input UserRegistrationInput, logg logger.Interface) (UserRegistrationOutput, error) {

	// Validate input data
	ok, err := govalidator.ValidateStruct(input)
	if !ok {
		logg.Info("Invalid input data for user registration")
		return UserRegistrationOutput{}, err
	}

	// Create a new user entity
	user, err := entity.NewUser(
		input.Username,
		input.Email,
		input.Password,
		entity.WithFullName(input.FullName),
		entity.WithBirthdate(input.Birthdate),
		entity.WithGender(input.Gender),
		entity.WithPhone(input.Phone),
	)
	if err != nil {
		logg.Error("Failed to create new user entity", "error", err)
		return UserRegistrationOutput{}, err
	}

	err = uc.UserRepository.Save(ctx, user)
	if err != nil {
		logg.Error("Failed to save user to the database", "error", err)
		return UserRegistrationOutput{}, err
	}

	return UserRegistrationOutput{}, nil
}
