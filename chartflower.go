package main

import (
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	csvFiles := chooseCsvFiles()
	table := makeSQLTable(csvFiles)
	chooseColumns(table)
}
