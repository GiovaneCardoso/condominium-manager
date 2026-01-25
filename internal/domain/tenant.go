package domain

import "time"

type Tenant struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Domain       string `json:"domain"`
	LogoUrl      string `json:"logo_url"`
	PrimaryColor string `json:"primary_color"`
	Status       string `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}
