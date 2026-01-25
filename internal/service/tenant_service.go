package service

import (
	"errors"
	"gerenciador-condominio/internal/domain"
	"gerenciador-condominio/internal/repository"
	"time"

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

func (s *TenantService) Create(t domain.Tenant) (*domain.Tenant, error) {
	if t.Name == "" {
		return nil, errors.New("Name is a required field")
	}
	if t.Domain == "" {
		return nil, errors.New("Domain is a required field")
	}
	_, err := s.repo.FindByDomainName(t.Domain)
	if err == nil {
		return nil, errors.New("This domain already exists")
	}
	t.ID = uuid.NewString()
	t.Status = "active"
	t.CreatedAt = time.Now()
	err = s.repo.Create(&t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
func (s *TenantService) List() ([]domain.Tenant, error) {
	return s.repo.List()
}
func (s *TenantService) Update(id string, t repository.TenantUpdate) (*domain.Tenant, error) {
	return s.repo.Update(id, t)
}
func (s *TenantService) Inactivate(id string) error {
	return s.repo.Inactivate(id)
}
