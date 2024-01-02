package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/kirimatt/http"
	"github.com/kirimatt/pg"
)

func main() {
	urlsSet, err := http.GetUrlsFromSearch("https://adilet.zan.kz/rus/search/docs/sort_desc=true&sort_field=dl")
	if err != nil {
		log.Fatalf("Failed to get urls: %s", err)
	}

	os.Setenv("DATABASE_URL", "postgres://user:password@localhost:5437/postgres")

	postgres, err := pg.NewPG(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to create postgres connection: %s", err)
	}

	urlsSet.Do(func(i interface{}) {
		url := fmt.Sprintf("https://adilet.zan.kz/rus/docs/%s/download/docx", i)
		fmt.Println(url)
		res, err := http.GetDataFromUrl(url)
		err = postgres.InsertGovermentAddon(context.Background(), url, *res)
		if err != nil {
			log.Fatalf("Failed to insert goverment addon: %s", err)
		}
	})
}
