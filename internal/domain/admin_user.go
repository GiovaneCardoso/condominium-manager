package domain

import "time"

type AdminUser struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Resources []string  `json:"resources"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
