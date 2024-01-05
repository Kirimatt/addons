package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (pg *postgres) InsertGovermentAddon(ctx context.Context, url string, status string, text string) error {
	query := `insert into goverment_addons (url, status, value) values (@url, @status, @text) on conflict do nothing`
	args := pgx.NamedArgs{
		"url":    url,
		"status": status,
		"text":   text,
	}
	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (pg *postgres) GetExistingGovermentAddons(ctx context.Context, urls []string, hash string) error {
	query := `insert into goverment_addons (url, status, value) values (@urls, @status, @text) on conflict do nothing`
	args := pgx.NamedArgs{
		"url":    url,
		"status": status,
		"text":   text,
	}
	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}
