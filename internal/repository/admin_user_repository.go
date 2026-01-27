package repository

import "gerenciador-condominio/internal/domain"

type AdminUserUpdate struct {
	Name      *string   `json:"name"`
	Email     *string   `json:"email"`
	Password  *string   `json:"password"`
	Resources *[]string `json:"resources"`
	Status    *string   `json:"status"`
}
type AdminUserRepository interface {
	Create(user *domain.AdminUser) error
	List() ([]domain.AdminUser, error)
	FindById(id string) (*domain.AdminUser, error)
	Update(id string, update AdminUserUpdate) (*domain.AdminUser, error)
	Inactivate(id string) error
}
