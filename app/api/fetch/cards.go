package fetch

import (
	"context"
	"database/sql"
	"main/api/models"

	"github.com/lib/pq"
)

type PostGresCardStore struct {
	db *sql.DB
}

func (storage *PostGresCardStore) Create(ctx context.Context, model models.Card) error {
	query := `INSERT INOT cards (scryfall_id, name, oracle_text, mana_cost, set, types)
			  VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := storage.db.QueryRowContext(
		ctx,
		query,
		model.ScryfallID,
		model.Name,
		model.OracleText,
		model.ManaCost,
		model.Set,
		pq.Array(model.Types),
	).Scan(
		&model.ScryfallID,
		&model.Name,
		&model.OracleText,
		&model.ManaCost,
		&model.Set,
	)
	if err != nil {
		return err
	}

	return nil
}
