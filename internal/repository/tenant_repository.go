package repository

import "gerenciador-condominio/internal/domain"

type TenantUpdate struct {
	Name         *string
	Domain       *string
	LogoURL      *string
	PrimaryColor *string
	Status       *string
}

type TenantRepository interface {
	Create(tentant *domain.Tenant) error
	FindById(id string) (*domain.Tenant, error)
	FindByDomainName(name string) (*domain.Tenant, error)
	Update(id string, update TenantUpdate) (*domain.Tenant, error)
	List() []domain.Tenant
}
