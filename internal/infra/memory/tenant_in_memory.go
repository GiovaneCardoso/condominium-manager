package memory

import (
	"errors"
	"sync"

	"gerenciador-condominio/internal/domain"
	"gerenciador-condominio/internal/repository"
)

type TenantInMemory struct {
	tenants map[string]domain.Tenant
	mu      sync.RWMutex
}

func NewTenantInMemory() *TenantInMemory {
	return &TenantInMemory{
		tenants: make(map[string]domain.Tenant),
	}
}

func (r *TenantInMemory) Create(tenant *domain.Tenant) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tenants[tenant.ID] = *tenant
	return nil
}

func (r *TenantInMemory) FindById(id string) (*domain.Tenant, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if tenant, exists := r.tenants[id]; exists {
		return &tenant, nil
	}
	return nil, errors.New("tenant not found")
}

func (r *TenantInMemory) FindByDomainName(name string) (*domain.Tenant, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, tenant := range r.tenants {
		if tenant.Domain == name {
			return &tenant, nil
		}
	}
	return nil, errors.New("tenant not found")
}

func (r *TenantInMemory) Update(id string, update repository.TenantUpdate) (*domain.Tenant, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	tenant, exists := r.tenants[id]
	if !exists {
		return nil, errors.New("tenant not found")
	}

	if update.Name != nil {
		tenant.Name = *update.Name
	}
	if update.Domain != nil {
		tenant.Domain = *update.Domain
	}
	if update.LogoURL != nil {
		tenant.LogoUrl = *update.LogoURL
	}
	if update.PrimaryColor != nil {
		tenant.PrimaryColor = *update.PrimaryColor
	}

	r.tenants[id] = tenant
	return &tenant, nil
}

func (r *TenantInMemory) Inactivate(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tenant, exists := r.tenants[id]
	if !exists {
		return errors.New("tenant not found")
	}
	tenant.Status = "inactive"
	r.tenants[id] = tenant
	return nil
}

func (r *TenantInMemory) List() ([]domain.Tenant, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tenants := make([]domain.Tenant, 0, len(r.tenants))
	for _, tenant := range r.tenants {
		tenants = append(tenants, tenant)
	}
	return tenants, nil
}
