package memory

import (
	"errors"
	"gerenciador-condominio/internal/domain"
	"gerenciador-condominio/internal/repository"
	"sync"
)

type AdminUserInMemory struct {
	users map[string]domain.AdminUser
	mu    sync.RWMutex
}

func NewAdminUserInMemory() *AdminUserInMemory {
	return &AdminUserInMemory{
		users: make(map[string]domain.AdminUser),
	}
}
func (r *AdminUserInMemory) Create(user *domain.AdminUser) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.ID] = *user
	return nil
}
func (r *AdminUserInMemory) FindById(id string) (*domain.AdminUser, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if user, exists := r.users[id]; exists {
		return &user, nil
	}
	return nil, errors.New("User not found")
}
func (r *AdminUserInMemory) Inactivate(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	user, exists := r.users[id]
	if !exists {
		return errors.New("User not found")
	}
	user.Status = "inactive"
	r.users[id] = user
	return nil
}
func (r *AdminUserInMemory) List() ([]domain.AdminUser, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	users := make([]domain.AdminUser, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}
func (r *AdminUserInMemory) Update(id string, update repository.AdminUserUpdate) (*domain.AdminUser, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("User not found")
	}
	if update.Email != nil {
		user.Email = *update.Email
	}
	if update.Name != nil {
		user.Name = *update.Name
	}
	if update.Email != nil {
		user.Email = *update.Email
	}
	if update.Resources != nil {
		user.Resources = *update.Resources
	}
	r.users[id] = user
	return &user, nil

}
