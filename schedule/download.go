package schedule

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/kirimatt/http"
	"github.com/kirimatt/pg"
)

func ScheduleDownloadingAddons() error {
	go func() {
		for {
			fmt.Println("Started scheduled task")
			err := scheduledDownloadingAddons()
			if err != nil {
				return
			}
			fmt.Println("Ended scheduled task")
			time.Sleep(100 * time.Second)
		}
	}()

	return nil
}

func scheduledDownloadingAddons() error {
	orderToUrlDescription, err := http.GetUrlsFromSearch(os.Getenv("SEARCH_URL"), os.Getenv("PLACEHOLDER_URL"))
	if err != nil {
		return fmt.Errorf("Failed to get urls: %w", err)
	}

	postgres, err := pg.NewPG(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return fmt.Errorf("Failed to create postgres connection: %w", err)
	}

	currentTime := time.Now()
	for orderId, urlDescription := range *orderToUrlDescription {
		processUrl(*postgres, orderId, urlDescription, currentTime)
		if err != nil {
			return fmt.Errorf("Failed on processing goverment addon: %w", err)
		}
	}

	return nil
}

func processUrl(postgres pg.Postgres, orderId int, urlDescription http.UrlDescription, time time.Time) error {
	fmt.Printf("Adding url: %s \n", urlDescription)
	res, finalUrl, err := http.GetDataFromUrl(urlDescription.Url)
	if err != nil {
		return fmt.Errorf("Failed on getting goverment addon: %w", err)
	}

	addon, err := postgres.GetExistingGovermentAddons(context.Background(), *finalUrl)
	if err != nil && err != pg.AddonDoesNotExist {
		return fmt.Errorf("Failed on checking existence of goverment addon: %w", err)
	}

	if addon != nil && addon.Status == urlDescription.Status {
		fmt.Printf("Url: %s, with status: %s already exists \n", urlDescription, urlDescription.Status)
		return nil
	}

	govermentAddon := pg.GovermentAddon{Url: *finalUrl, Status: urlDescription.Status, Text: *res, OrderId: orderId, CreatedAt: &time}
	err = postgres.InsertGovermentAddon(context.Background(), &govermentAddon)
	if err != nil {
		return fmt.Errorf("Failed to insert goverment addon: %w", err)
	}
	fmt.Printf("Added url: %s, with status: %s \n", urlDescription, urlDescription.Status)

	return nil
}
