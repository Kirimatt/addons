package pg

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type GovermentAddon struct {
	Url       string
	Status    string
	Text      string
	OrderId   int
	CreatedAt *time.Time
}

var (
	AddonDoesNotExist = errors.New("no rows in result set")
)

func (pg *Postgres) InsertGovermentAddon(ctx context.Context, govermentAddon *GovermentAddon) error {
	query := `insert into goverment_addons (url, status, value, order_id, created_at) values (@url, @status, @text, @order_id, @created_at) on conflict do nothing`
	args := pgx.NamedArgs{
		"url":        govermentAddon.Url,
		"status":     govermentAddon.Status,
		"text":       govermentAddon.Text,
		"order_id":   govermentAddon.OrderId,
		"created_at": govermentAddon.CreatedAt,
	}
	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (pg *Postgres) GetExistingGovermentAddons(ctx context.Context, url string) (addon *GovermentAddon, err error) {
	fmt.Printf("Searching by url: %s \n", url)
	query := `select url, status from goverment_addons where url = @url`
	args := pgx.NamedArgs{
		"url": url,
	}

	var result GovermentAddon
	err = pg.db.QueryRow(context.Background(), query, args).Scan(&result.Url, &result.Status)
	if err == pgx.ErrNoRows {
		fmt.Printf("Not found by url: %s \n", url)
		return nil, AddonDoesNotExist
	}
	if err != nil {
		return nil, fmt.Errorf("Error by getting rows: %s %w \n", url, err)
	}
	fmt.Printf("Found: %+v \n", result)

	return &result, nil
}
