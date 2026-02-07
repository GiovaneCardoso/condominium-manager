package postgres

import (
	"database/sql"
	"errors"

	"gerenciador-condominio/internal/domain"
	"gerenciador-condominio/internal/repository"
)

type AdminUserPostgres struct {
	db *sql.DB
}

func NewAdminUserPostgres(db *sql.DB) *AdminUserPostgres {
	return &AdminUserPostgres{db: db}
}

func (r *AdminUserPostgres) Create(user *domain.AdminUser) error {
	query := `
		INSERT INTO admin_users (
			id,
			name,
			email,
			password,
			status
		) VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(
		query,
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.Status,
	)

	return err
}
func (r *AdminUserPostgres) FindByEmail(email string) (*domain.AdminUser, error) {
	query := `
		SELECT id, name, email, status
		FROM admin_users
		WHERE email = $1
	`

	var user domain.AdminUser

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Status,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (r *AdminUserPostgres) List() ([]domain.AdminUser, error) {
	query := `
		SELECT id, name, email, status
		FROM admin_users
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]domain.AdminUser, 0)

	for rows.Next() {
		var user domain.AdminUser

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Status,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
func (r *AdminUserPostgres) FindById(id string) (*domain.AdminUser, error) {
	query := `
		SELECT id, name, email, status
		FROM admin_users
		WHERE id = $1
	`

	var user domain.AdminUser

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Status,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AdminUserPostgres) Update(id string, update repository.AdminUserUpdate) (*domain.AdminUser, error) {
	query := `
		UPDATE admin_users
		SET name = COALESCE($1, name),
		    email = COALESCE($2, email),
		    password = COALESCE($3, password),
		    status = COALESCE($4, status)
		WHERE id = $5
		RETURNING id, name, email, status
	`

	var user domain.AdminUser

	err := r.db.QueryRow(
		query,
		update.Name,
		update.Email,
		update.Password,
		update.Status,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Status,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AdminUserPostgres) Inactivate(id string) error {
	query := `
		UPDATE admin_users
		SET status = 'inactive'
		WHERE id = $1
	`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
