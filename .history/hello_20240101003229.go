package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	db, err := sql.Open("postgres",
		"user:password@tcp(127.0.0.1:5437)/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
