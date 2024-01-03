package main

import (
	"fmt"
	"os"

	"github.com/kirimatt/schedule"
)

var (
	searchUrl      = "https://adilet.zan.kz/rus/search/docs/sort_desc=true&sort_field=dl"
	databaseUrl    = os.Getenv("DATABASE_URL")
	placeholderUrl = "https://adilet.zan.kz/rus/docs/%s/download/docx"
)

func main() {
	err := schedule.ScheduleDownloadingDocs(
		searchUrl,
		databaseUrl,
		placeholderUrl,
	)
	if err != nil {
		fmt.Errorf("An error occured when scheduling task: %w", err)
	}

	os.Setenv("DATABASE_URL", "postgres://user:password@localhost:5437/postgres")
}