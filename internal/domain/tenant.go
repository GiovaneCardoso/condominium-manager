package domain

import "time"

type Tenant struct {
	ID           string
	Name         string
	Domain       string
	LogoUrl      string
	PrimaryColor string
	Status       string
	CreatedAt    time.Time
}
