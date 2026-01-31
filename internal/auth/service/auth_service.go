package authService

import (
	"errors"
	"gerenciador-condominio/internal/repository"
	"gerenciador-condominio/internal/service"
)

type AuthService struct {
	repo         repository.AdminUserRepository
	tokenService *TokenService
}

func NewAuthService(repo repository.AdminUserRepository, tokenService *TokenService) *AuthService {
	return &AuthService{repo: repo, tokenService: tokenService}
}

func (s *AuthService) Authenticate(email, password string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("email and password are required")
	}

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !service.CheckPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	return s.tokenService.GenerateAdminToken(user.ID)
}
