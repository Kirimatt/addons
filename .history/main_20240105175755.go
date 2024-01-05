package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/kirimatt/pg"
	"github.com/kirimatt/schedule"
)

var (
	searchUrl      = os.Getenv("SEARCH_URL")
	databaseUrl    = os.Getenv("DATABASE_URL")
	placeholderUrl = os.Getenv("PLACEHOLDER_URL")
)

func main() {
	if databaseUrl == "" {
		os.Setenv("DATABASE_URL", "postgres://user:password@localhost:5437/postgres")
	}
	if searchUrl == "" {
		os.Setenv("SEARCH_URL", "https://adilet.zan.kz/rus/search/docs/sort_desc=true&sort_field=dl")
	}
	if placeholderUrl == "" {
		os.Setenv("PLACEHOLDER_URL", "https://adilet.zan.kz/rus/docs/%s/download/docx")
	}

	err := schedule.ScheduleDownloadingAddons()
	if err != nil {
		fmt.Println("An error occured when scheduling task: %w", err)
	}

	pg, err := pg.NewPG(context.Background(), os.Getenv("DATABASE_URL"))
	pg.GetAddonsToProcess()

	time.Sleep(40 * time.Second)
}
