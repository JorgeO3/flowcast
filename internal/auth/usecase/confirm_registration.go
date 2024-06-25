// Package usecase contains the implementation of the user registration confirmation use case.
package usecase

import (
	"context"

	"gitlab.com/JorgeO3/flowcast/configs"
	"gitlab.com/JorgeO3/flowcast/internal/auth/repository"
	"gitlab.com/JorgeO3/flowcast/pkg/logger"
)

// ConfirmRegInput represents the input data required for user registration confirmation.
type ConfirmRegInput struct {
	UserID int    `json:"user_id" valid:"required"`
	Token  string `json:"token" valid:"required"`
}

// ConfirmRegOutput represents the output data returned after user registration confirmation.
type ConfirmRegOutput struct {
	Success bool `json:"success"`
}

// ConfirmRegUC handles the user registration confirmation logic.
type ConfirmRegUC struct {
	UserRepo repository.UserRepo
	Logg     logger.Interface
}

// NewConfirmRegUC creates a new instance of ConfirmRegUC.
func NewConfirmRegUC(userRepo repository.UserRepo, logg logger.Interface) *ConfirmRegUC {
	return &ConfirmRegUC{
		UserRepo: userRepo,
		Logg:     logg,
	}
}

// Execute performs the user registration confirmation.
func (uc *ConfirmRegUC) Execute(ctx context.Context, input ConfirmRegInput, cfg *configs.AuthConfig) (ConfirmRegOutput, error) {
	// Implement the registration confirmation logic here...
	_, err := uc.UserRepo.FindByID(ctx, input.UserID)
	if err != nil {
		uc.Logg.Error("Failed to find user", "error", err)
		return ConfirmRegOutput{Success: false}, err
	}

	// Add logic to verify the token and update user status

	// Example:
	// if user != nil && user.ConfirmationToken == input.Token {
	// 	user.Status = "active"
	// 	err = uc.UserRepo.Update(context.Background(), user)
	// 	if err != nil {
	// 		return ConfirmRegOutput{Success: false}, err
	// 	}
	// 	return ConfirmRegOutput{Success: true}, nil
	// }

	return ConfirmRegOutput{Success: false}, nil
}
