package schedule

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/kirimatt/http"
	"github.com/kirimatt/pg"
)

func ScheduleDownloadingDocs(searchUrl string, databaseUrl string, placeholderUrl string) error {
	go func() {
		for {
			fmt.Println("Started scheduled task")
			err := scheduledTask(searchUrl, databaseUrl, placeholderUrl)
			if err != nil {
				return
			}
			fmt.Println("Ended scheduled task")
			time.Sleep(10 * time.Second)
		}
	}()

	return nil
}

func scheduledTask(searchUrl string, databaseUrl string, placeholderUrl string) error {
	docToStatus, unstructedData, err := http.GetUrlsFromSearch(searchUrl, placeholderUrl)
	if err != nil {
		return fmt.Errorf("Failed to get urls: %w", err)
	}

	postgres, err := pg.NewPG(context.Background(), databaseUrl)
	if err != nil {
		return fmt.Errorf("Failed to create postgres connection: %w", err)
	}

	sha := sha256.New()
	sha.Write([]byte(unstructedData))
	hash := sha.Sum(nil)

	for url, status := range *docToStatus {
		fmt.Printf("Adding url: %s \n", url)
		res, err := http.GetDataFromUrl(url)
		if err != nil {
			return fmt.Errorf("Failed on getting goverment addon: %w", err)
		}
		err = postgres.InsertGovermentAddon(context.Background(), url, status, *res, string(hash))
		if err != nil {
			return fmt.Errorf("Failed to insert goverment addon: %w", err)
		}
		fmt.Printf("Added url: %s, with status: %s \n", url, status)
	}

	return nil
}
