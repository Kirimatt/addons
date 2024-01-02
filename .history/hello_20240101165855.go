package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/unidoc/unioffice/document"

	"github.com/jackc/pgx/v5"
)

func main() {
	doc, err := document.Open("https://adilet.zan.kz/rus/docs/K950001000_/download/docx")
	if err != nil {
		log.Fatalf("Failed to open document: %s", err)
	}

	// Получение содержимого текстовых элементов из документа
	para := doc.Paragraphs()
	for _, p := range para {
		content := ""
		for _, run := range p.Runs() {
			content += run.Text()
		}
		fmt.Println(content)
	}
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
