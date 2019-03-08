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
	columns := chooseColumns(csvFiles)
	fmt.Println(columns)
}

func chooseColumns(csvFiles []string) []string {
	var allColumns []string
	for _, csvFile := range csvFiles {
		csv := convertCSVToArray(csvFile)
		columns := getColumns(csv)
		for _, column := range columns {
			if stringInSlice(column, allColumns) == false {
				allColumns = append(allColumns, column)
			}
		}
	}
	println()
	for i, column := range allColumns {
		fmt.Printf("%v. %v\n", i, column)
	}
	print("\nChoose columns:")
	input := getConsoleText()
	choices := strings.Fields(input)
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

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getColumns(csv [][]string) []string {
	return csv[0]
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
