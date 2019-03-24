package main

import (
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	data := csvToArray()
	chart := makeChart(data)
	saveFile(chart)
}
