package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	csv := selectCSV()
	sql := convertCSVToSQL(csv)
	fmt.Println(sql)
}
