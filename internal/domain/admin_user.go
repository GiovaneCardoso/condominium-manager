package domain

import "time"

type AdminUser struct {
	ID        string     `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Name      string     `json:"name"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	Resources []Resource `json:"resources"`
}

type Resource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
