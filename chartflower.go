package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	data := csvToArray()
	chart := makeChart(data)
	// makeFile(csv, chart)

	fmt.Println(chart)

}
