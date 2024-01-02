package main

import (
	"fmt"

	_ "github.com/jackc/pgx"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
