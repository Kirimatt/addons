package main

import (
	"fmt"
	"log"

	"github.com/kirimatt/http"
)

func main() {
	urls, err := http.GetUrlsFromSearch("https://adilet.zan.kz/rus/search/docs/sort_desc=true&sort_field=dl")
	if err != nil {
		log.Fatalf("Failed to get urls: %s", err)
	}
	fmt.Println(urls)

	// os.Setenv("DATABASE_URL", "postgres://user:password@localhost:5437/postgres")

	// url := "https://adilet.zan.kz/rus/docs/K950001000_/download/docx"
	// res, err := http.GetDataFromUrl(url)

	// postgres, err := pg.NewPG(context.Background(), os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatalf("Failed to create postgres connection: %s", err)
	// }

	// err = postgres.InsertGovermentAddon(context.Background(), url, *res)
	// if err != nil {
	// 	log.Fatalf("Failed to insert goverment addon: %s", err)
	// }
}
