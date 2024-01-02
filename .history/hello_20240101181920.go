package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"code.sajari.com/docconv/v2"

	"github.com/jackc/pgx/v5"
)

func main() {
	url := "https://adilet.zan.kz/rus/docs/K950001000_/download/docx"

	initResp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make GET request: %s", err)
	}
	defer initResp.Body.Close()

	finalURL := initResp.Request.URL.String()
	fmt.Println(finalURL)
	resp, err := http.Get(finalURL)
	if err != nil {
		log.Fatalf("Failed to make GET request: %s", err)
	}
	defer resp.Body.Close()

	fileData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %s", err)
	}

	tmpFile, err := os.CreateTemp("", "docxfile*.docx")
	if err != nil {
		log.Fatalf("Failed to create temporary file: %s", err)
	}

	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write(fileData)
	if err != nil {
		log.Fatalf("Failed to write to temporary file: %s", err)
	}

	res, err := docconv.ConvertPath(tmpFile.Name())
	if err != nil {
		log.Fatalf("Failed to convert temporary file: %s", err)
	}
	fmt.Println(res)
	// resp, err := http.Get("https://adilet.zan.kz/rus/docs/K950001000_/download/docx")
	// if err != nil {
	// 	// handle error
	// }
	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// fmt.Println(string(body))

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
