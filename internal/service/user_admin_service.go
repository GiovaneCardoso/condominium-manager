package service

import (
	"errors"
	"gerenciador-condominio/internal/domain"
	"gerenciador-condominio/internal/repository"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserAdminService struct {
	repo repository.AdminUserRepository
}

func NewUserAdminService(repo repository.AdminUserRepository) *UserAdminService {
	return &UserAdminService{
		repo: repo,
	}
}
func (s *UserAdminService) Create(u domain.AdminUser) (*domain.AdminUser, error) {
	if u.Email == "" {
		return nil, errors.New("Email is a required field")

	}
	if u.Name == "" {
		return nil, errors.New("Name is a required field")
	}
	if u.Password == "" {
		return nil, errors.New("Password is a required field")
	}
	u.ID = uuid.NewString()
	u.Status = "active"
	u.CreatedAt = time.Now()
	hash, err := HashPassword(u.Password)
	if err != nil {
		return nil, errors.New("Internal server error")
	}
	u.Password = hash
	err = s.repo.Create(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func (s *UserAdminService) List() ([]domain.AdminUser, error) {
	return s.repo.List()
}
func (s *UserAdminService) Update(id string, u repository.AdminUserUpdate) (*domain.AdminUser, error) {
	return s.repo.Update(id, u)
}
func (s *UserAdminService) Inactivate(id string) error {
	return s.repo.Inactivate(id)
}
func (s *UserAdminService) FindById(id string) (*domain.AdminUser, error) {
	return s.repo.FindById(id)
}
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	return string(hash), err
}
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
	return err == nil
}
