package controller

import (
	"net/http"

	"gitlab.com/JorgeO3/flowcast/internal/auth/usecase"
)

type AuthController struct {
	UserRegistrationUseCase    *usecase.UserRegistrationUseCase
	UserAuthenticationUseCase  *usecase.UserAuthenticationUseCase
	ConfirmRegistrationUseCase *usecase.ConfirmRegistrationUseCase
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {

}

func (c *AuthController) Authenticate(w http.ResponseWriter, r *http.Request) {

}

func (c *AuthController) ConfirmRegistration(w http.ResponseWriter, r *http.Request) {

}
