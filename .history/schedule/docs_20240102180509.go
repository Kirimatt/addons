package schedule

import (
	"context"
	"fmt"
	"time"

	"github.com/kirimatt/http"
	"github.com/kirimatt/pg"
	"github.com/madflojo/tasks"
)

func scheduleDownloadingDocs(searchUrl string, databaseUrl string, placeholderUrl string) error {
	// Start the Scheduler
	scheduler := tasks.New()
	defer scheduler.Stop()

	// Add a task
	id, err := scheduler.Add(&tasks.Task{
		Interval: time.Duration(30 * time.Second),
		TaskFunc: func() error {
			// Put your logic here
		},
	})
	fmt.Println("Started task with id: %s", id)
	if err != nil {
		// Do Stuff
	}

	return nil
}

func scheduledTask(searchUrl string, databaseUrl string, placeholderUrl string) error {
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
		fmt.Println("Adding url: %s", url)
		res, err := http.GetDataFromUrl(url)
		if err != nil {
			return fmt.Errorf("Failed on getting goverment addon: %w", err)
		}
		err = postgres.InsertGovermentAddon(context.Background(), url, value, *res)
		if err != nil {
			return fmt.Errorf("Failed to insert goverment addon: %w", err)
		}
		fmt.Println("Added url: %s, with status: %s", url, value)
	}

	return nil
}
