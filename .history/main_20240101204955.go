package main

import (
	"context"
	"log"
	"os"

	"github.com/kirimatt/http"
	"github.com/kirimatt/pg"
)

func main() {
	url := "https://adilet.zan.kz/rus/docs/K950001000_/download/docx"
	http.GetDataFromUrl(url)
	// resp, err := http.Get("https://adilet.zan.kz/rus/docs/K950001000_/download/docx")
	// if err != nil {
	// 	// handle error
	// }
	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// fmt.Println(string(body))

	os.Setenv("DATABASE_URL", "postgres://user:password@localhost:5437/postgres")
	// conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer conn.Close(context.Background())

	// _, err = conn.Exec(context.Background(), "insert into goverment_addons (url, value) values (:url, 'some text')", )
	// var greeting string
	// err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }
	postgres, err := pg.NewPG(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to create postgres connection: %s", err)
	}

	err = postgres.InsertGovermentAddon(context.Background(), url, res.Body)
	if err != nil {
		log.Fatalf("Failed to insert goverment addon: %s", err)
	}
}
