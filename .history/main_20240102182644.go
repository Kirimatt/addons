package main

import (
	"fmt"
	"os"

	"github.com/kirimatt/schedule"
)

var ()

func main() {
	err := schedule.ScheduleDownloadingDocs(
		"https://adilet.zan.kz/rus/search/docs/sort_desc=true&sort_field=dl",
		os.Getenv("DATABASE_URL"),
		"https://adilet.zan.kz/rus/docs/%s/download/docx",
	)
	if err != nil {
		fmt.Errorf("An error occured when scheduling task: %w", err)
	}

	os.Setenv("DATABASE_URL", "postgres://user:password@localhost:5437/postgres")
}
