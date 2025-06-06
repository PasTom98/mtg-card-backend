package fetch

import (
	"context"
	"database/sql"
	"main/api/models"
)

type Storage struct {
	Cards interface {
		Create(ctx context.Context, model models.Card) error
	}
	Users interface {
		Create(ctx context.Context, user models.User) error
	}
}

func NewPostGresStorage(db *sql.DB) Storage {
	return Storage{
		Cards: &PostGresCardStore{db: db},
		Users: &PostgresUsersStore{db: db},
	}
}
