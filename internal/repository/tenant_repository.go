package repository

import "gerenciador-condominio/internal/domain"

type TenantUpdate struct {
	Name         *string `json:"name"`
	Domain       *string `json:"domain"`
	LogoURL      *string `json:"logo_url"`
	PrimaryColor *string `json:"primary_color"`
}

type TenantRepository interface {
	Create(tentant *domain.Tenant) error
	FindById(id string) (*domain.Tenant, error)
	FindByDomainName(name string) (*domain.Tenant, error)
	Update(id string, update TenantUpdate) (*domain.Tenant, error)
	List() ([]domain.Tenant, error)
	Inactivate(id string) error
}
