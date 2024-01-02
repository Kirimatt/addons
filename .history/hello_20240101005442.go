package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

func readHtmlFromFile(fileName string) (string, error) {

	bs, err := io.ReadAll(fileName)

	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func main() {
	resp, err := http.Get("https://adilet.zan.kz/rus/docs/K950001000_")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	os.Setenv("DATABASE_URL", "postgres://user:password@localhost:5437/postgres")
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}