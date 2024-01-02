package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
