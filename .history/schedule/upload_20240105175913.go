package schedule

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/kirimatt/pg"
)

func ScheduleUploadingAddons(searchUrl string, databaseUrl string, placeholderUrl string) error {
	go func() {
		for {
			fmt.Println("Started scheduled task")
			err := scheduledUploadingAddons()
			if err != nil {
				return
			}
			fmt.Println("Ended scheduled task")
			time.Sleep(100 * time.Second)
		}
	}()

	return nil
}

func scheduledUploadingAddons() error {
	postgres, err := pg.NewPG(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return fmt.Errorf("Failed to create postgres connection: %w", err)
	}

	fmt.Printf("postgres: %s", postgres)

	return nil
}
