package main

import (
	"fmt"
	"os"
	"time"

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

	err := schedule.ScheduleDownloadingAddons(
		placeholderUrl,
	)
	if err != nil {
		fmt.Println("An error occured when scheduling task: %w", err)
	}

	time.Sleep(40 * time.Second)
}
