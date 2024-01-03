package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kirimatt/schedule"
)

var (
	searchUrl      = "https://adilet.zan.kz/rus/search/docs/sort_desc=true&sort_field=dl"
	databaseUrl    = os.Getenv("DATABASE_URL")
	placeholderUrl = "https://adilet.zan.kz/rus/docs/%s/download/docx"
)

func main() {
	if databaseUrl == "" {
		databaseUrl = "postgres://user:password@localhost:5437/postgres"
	}

	err := schedule.ScheduleDownloadingDocs(
		searchUrl,
		databaseUrl,
		placeholderUrl,
	)
	if err != nil {
		fmt.Println("An error occured when scheduling task: %w", err)
	}

	time.Sleep(40 * time.Second)
}
