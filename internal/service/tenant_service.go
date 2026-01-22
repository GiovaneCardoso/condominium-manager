package service

import (
	"errors"
	"gerenciador-condominio/internal/domain"
	"gerenciador-condominio/internal/repository"

	"github.com/google/uuid"
)

type TenantService struct {
	repo repository.TenantRepository
}

func NewTenantService(repo repository.TenantRepository) *TenantService {
	return &TenantService{
		repo: repo,
	}
}

func (s *TenantService) Create(t domain.Tenant) error {
	if t.Name == "" {
		return errors.New("Name is a required field")
	}
	if t.Domain == "" {
		return errors.New("Domain is a required field")
	}
	_, err := s.repo.FindByDomainName(t.Domain)
	if err == nil {
		return errors.New("This domain already exists")
	}
	t.ID = uuid.NewString()
	return s.repo.Create(&t)
}
func (s *TenantService) List() ([]domain.Tenant, error) {

	return s.repo.List(), nil
}
