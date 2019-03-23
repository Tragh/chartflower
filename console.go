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

func csvToArray() [][]string {
	var selectedCSV string

	csvFilenames := getCSVFilenames()
	for i, csvFilename := range csvFilenames {
		csvTrimmed := strings.TrimSuffix(csvFilename, ".csv")
		fmt.Println(strconv.Itoa(i) + ". " + csvTrimmed)
	}

	fmt.Print("\nChoose a CSV: ")
	choice := getConsoleText()

	for i, csvFilename := range csvFilenames {
		csvTrimmed := strings.TrimSuffix(csvFilename, ".csv")

		if choice == csvFilename {
			selectedCSV = csvFilename
		} else if choice == csvTrimmed {
			selectedCSV = csvFilename
		} else if choice == strconv.Itoa(i) {
			selectedCSV = csvFilename
		}
	}

	data := convertCSVToArray(selectedCSV)
	return data
}

func convertCSVToArray(filename string) [][]string {
	file, _ := os.Open("./csv/" + filename)
	defer file.Close()
	reader := csv.NewReader(file)
	rows, _ := reader.ReadAll()
	return rows
}

func getCSVFilenames() []string {
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
