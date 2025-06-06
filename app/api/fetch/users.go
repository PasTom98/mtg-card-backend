package fetch

import (
	"context"
	"database/sql"
	"main/api/models"
)

type UserStore interface {
	GetByID(id int) (*User, error)
}

type PostgresUsersStore struct {
	db *sql.DB
}

func (storage *PostgresUsersStore) Create(ctx context.Context, model models.User) error {
	query := `INSERT INTO users (username, password, email) VALUES($1, $2, $3) 
			  RETURNING id, created_at`

	err := storage.db.QueryRowContext(
		ctx,
		query,
		model.Username,
		model.Email,
		model.Password,
	).Scan(
		&model.Username,
		&model.Email,
		&model.Password,
	)
	if err != nil {
		return err
	}

	return nil
}

type User struct {
	ID   int
	Name string
}

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}
