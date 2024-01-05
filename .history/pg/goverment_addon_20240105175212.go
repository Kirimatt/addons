package pg

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

type GovermentAddon struct {
	Id        int64
	Url       string
	Status    string
	Text      string
	OrderId   int
	CreatedAt *time.Time
}

var (
	AddonDoesNotExist = errors.New("no rows in result set")
	NoAddonsToProcess = errors.New("no rows for processing")
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

func (pg *Postgres) GetAddonsToProcess() (addons []*GovermentAddon, err error) {
	fmt.Printf("Searching addons for processing \n")
	query := `select id, url, status, value from goverment_addons where is_processed order by created_at, order_id`

	var results []GovermentAddon
	rows, err := pg.db.Query(context.Background(), query)
	if err == pgx.ErrNoRows {
		fmt.Printf("Not found for processing \n")
		return nil, NoAddonsToProcess
	}
	defer rows.Close()

	for rows.Next() {
		var row GovermentAddon
		err := rows.Scan(&row.Id, &row.Url, &row.Status, &row.Text)
		if err != nil {
			return nil, fmt.Errorf("Error by getting rows for processing: %w \n", err)
		}
		rowSlice = append(rowSlice, row)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(rowSlice)
	if err != nil {
		return nil, fmt.Errorf("Error by getting rows: %s %w \n", url, err)
	}
	fmt.Printf("Found: %+v \n", results)

	return &results, nil
}
