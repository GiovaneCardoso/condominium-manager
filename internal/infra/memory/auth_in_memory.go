package memory

import (
	"errors"
	"gerenciador-condominio/internal/repository"
)

type AuthInMemory struct {
	adminUserRepo repository.AdminUserRepository
}

func NewAuthInMemory(adminUserRepo repository.AdminUserRepository) *AuthInMemory {
	return &AuthInMemory{
		adminUserRepo: adminUserRepo,
	}
}

func (a *AuthInMemory) Authenticate(email, password string) (string, error) {
	user, err := a.adminUserRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if user.Status != "active" {
		return "", errors.New("user is not active")
	}

	return user.ID, nil
}
