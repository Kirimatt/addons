package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (pg *postgres) InsertGovermentAddon(ctx context.Context, url string, text string) error {
	query := `insert into goverment_addons (url, value) values (@url, @text)`
	args := pgx.NamedArgs{
		"url":  url,
		"text": text,
	}
	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}
