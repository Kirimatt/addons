package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (pg *postgres) InsertGovermentAddon(ctx context.Context) error {
	query := `INSERT INTO users (name, email) VALUES (@userName, @userEmail)`
	args := pgx.NamedArgs{
		"userName":  "Bobby",
		"userEmail": "bobby@donchev.is",
	}
	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}
