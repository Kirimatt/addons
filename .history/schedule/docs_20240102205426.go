package schedule

import (
	"context"
	"fmt"
	"time"

	"github.com/kirimatt/http"
	"github.com/kirimatt/pg"
)

func ScheduleDownloadingDocs(searchUrl string, databaseUrl string, placeholderUrl string) error {
	go func() {
		for true {
			err := ScheduledTask(searchUrl, databaseUrl, placeholderUrl)
			if err != nil {
				return
			}
			time.Sleep(10 * time.Second)
		}
	}()

	return nil
}

func ScheduledTask(searchUrl string, databaseUrl string, placeholderUrl string) error {
	docToStatus, err := http.GetUrlsFromSearch(searchUrl)
	if err != nil {
		return fmt.Errorf("Failed to get urls: %w", err)
	}

	postgres, err := pg.NewPG(context.Background(), databaseUrl)
	if err != nil {
		return fmt.Errorf("Failed to create postgres connection: %w", err)
	}

	for key, value := range *docToStatus {
		url := fmt.Sprintf(placeholderUrl, key)
		fmt.Printf("Adding url: %s", url)
		res, err := http.GetDataFromUrl(url)
		if err != nil {
			return fmt.Errorf("Failed on getting goverment addon: %w", err)
		}
		err = postgres.InsertGovermentAddon(context.Background(), url, value, *res)
		if err != nil {
			return fmt.Errorf("Failed to insert goverment addon: %w", err)
		}
		fmt.Printf("Added url: %s, with status: %s", url, value)
	}

	return nil
}
