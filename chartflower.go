package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	csvFiles := chooseCsvFiles()
	chosenColumns := chooseColumns(csvFiles)
	getColumnData(chosenColumns, csvFiles)
}

func getColumnData(chosenColumns []string, csvFiles []string) [][]string {
	var columns []string
	var columnIndexes []string
	var columnData [][]string
	for csvIndex, csvFile := range csvFiles {
		csv := convertCSVToArray(csvFile)
		for _, row := range csv {
			for columnIndex, column := range row {
				for _, chosenColumn := range chosenColumns {
					if chosenColumn == column {
						if stringInSlice(chosenColumn, columns) == false {
							columns = append(columns, chosenColumn)
							columnIndexes = append(columnIndexes, strconv.Itoa(csvIndex)+" "+strconv.Itoa(columnIndex))
						}
					}
				}
			}
		}
	}
	for _, columnIndex := range columnIndexes {
		indexStrings := strings.Fields(columnIndex)
		csvIndex, _ := strconv.Atoi(indexStrings[0])
		columnIndex, _ := strconv.Atoi(indexStrings[1])
		for i, csvFile := range csvFiles {
			if i == csvIndex {
				var columnRows []string
				csv := convertCSVToArray(csvFile)
				for _, row := range csv {
					columnRows = append(columnRows, row[columnIndex])
				}
				columnData = append(columnData, columnRows)
			}
		}
	}
	return columnData
}

func chooseColumns(csvFiles []string) []string {
	allColumns := getAllColumns(csvFiles)
	println()
	for i, column := range allColumns {
		fmt.Printf("%v. %v\n", i, column)
	}
	print("\nChoose columns:")
	input := getConsoleText()
	choices := strings.Fields(input)
	chosenColumns := getChosenColumns(choices, allColumns)
	return chosenColumns
}

func getChosenColumns(choices []string, allColumns []string) []string {
	var chosenColumns []string
	for i, column := range allColumns {
		for _, choice := range choices {
			if choice == column {
				chosenColumns = append(chosenColumns, column)
			} else if choice == strconv.Itoa(i) {
				chosenColumns = append(chosenColumns, column)
			}
		}
	}
	return chosenColumns
}

func getAllColumns(csvFiles []string) []string {
	var allColumns []string
	for _, csvFile := range csvFiles {
		csv := convertCSVToArray(csvFile)
		columnTitles := csv[0]
		for _, column := range columnTitles {
			if stringInSlice(column, allColumns) == false {
				allColumns = append(allColumns, column)
			}
		}
	}
	return allColumns
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func convertCSVToArray(filename string) [][]string {
	filename = "./csv/" + filename
	file, _ := os.Open(filename)
	defer file.Close()
	reader := csv.NewReader(file)
	rows, _ := reader.ReadAll()
	return rows
}

func chooseCsvFiles() []string {
	csvFiles := getCsvFilesinFolder()
	println()
	for index, csvFiles := range csvFiles {
		csvTrimmed := strings.TrimSuffix(csvFiles, ".csv")
		fmt.Printf("%v. %v\n", index, csvTrimmed)
	}
	print("\nChoose CSVs:")
	input := getConsoleText()
	choices := strings.Fields(input)
	var chosenCsvFiles []string
	for index, csvFile := range csvFiles {
		csvTrimmed := strings.TrimSuffix(csvFile, ".csv")
		for _, choice := range choices {
			if choice == csvFile {
				chosenCsvFiles = append(chosenCsvFiles, csvFile)
			} else if choice == csvTrimmed {
				chosenCsvFiles = append(chosenCsvFiles, csvFile)
			} else if choice == strconv.Itoa(index) {
				chosenCsvFiles = append(chosenCsvFiles, csvFile)
			}
		}
	}
	return chosenCsvFiles
}

func getCsvFilesinFolder() []string {
	var csvFiles []string
	filesInFolder, _ := ioutil.ReadDir("./csv/")
	for _, fileInFolder := range filesInFolder {
		if strings.HasSuffix(fileInFolder.Name(), "csv") {
			csvFiles = append(csvFiles, fileInFolder.Name())
		}
	}
	return csvFiles
}

func getConsoleText() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = removeNewLine(text)
	typeQToQuit(text)
	return text
}

func removeNewLine(input string) string {
	switch os := runtime.GOOS; os {
	case "windows":
		newLineRemoved := strings.TrimSuffix(input, "\r\n")
		return newLineRemoved
	default:
		newLineRemoved := strings.TrimSuffix(input, "\n")
		return newLineRemoved
	}
}

func typeQToQuit(input string) {
	if input == "q" {
		os.Exit(0)
	}
}
