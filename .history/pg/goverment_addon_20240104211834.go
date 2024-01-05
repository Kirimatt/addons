package pg

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type govermentAddon struct {
	url    string
	status string
}

func (addon govermentAddon) Status() string {
	return addon.status
}

func (addon govermentAddon) Url() string {
	return addon.url
}

var (
	AddonDoesNotExist = errors.New("no rows in result set")
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

func (pg *postgres) GetExistingGovermentAddons(ctx context.Context, url string) (addon *govermentAddon, err error) {
	fmt.Printf("Searching by url: %s \n", url)
	query := `select url, status from goverment_addons where url = @url`
	args := pgx.NamedArgs{
		"url": url,
	}

	var result govermentAddon
	err = pg.db.QueryRow(context.Background(), query, args).Scan(&result.url, &result.status)
	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("Not found by url: %s %w \n", url, AddonDoesNotExist)
	}
	if err != nil {
		return nil, fmt.Errorf("Error by getting rows: %s %w \n", url, err)
	}
	fmt.Printf("Found: %s \n", result)

	return &result, nil
}
