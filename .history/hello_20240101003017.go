package main

import (
	"database/sql"
	"fmt"
	"log"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	db, err := sql.Open("postgresql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
