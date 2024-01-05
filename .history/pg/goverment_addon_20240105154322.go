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
	text   string
	order  int
}

type govermentAddonDto struct {
	url    string
	status string
}

func (addon govermentAddonDto) Status() string {
	return addon.status
}

func (addon govermentAddonDto) Url() string {
	return addon.url
}

var (
	AddonDoesNotExist = errors.New("no rows in result set")
)

func (pg *postgres) InsertGovermentAddon(ctx context.Context, govermentAddon govermentAddon) error {
	query := `insert into goverment_addons (url, status, value) values (@url, @status, @text) on conflict do nothing`
	args := pgx.NamedArgs{
		"url":    govermentAddon.url,
		"status": govermentAddon.status,
		"text":   govermentAddon.text,
		"order":  govermentAddon.order,
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
		fmt.Printf("Not found by url: %s \n", url)
		return nil, AddonDoesNotExist
	}
	if err != nil {
		return nil, fmt.Errorf("Error by getting rows: %s %w \n", url, err)
	}
	fmt.Printf("Found: %s \n", result)

	return &result, nil
}
