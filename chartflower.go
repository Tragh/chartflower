package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	chooseCsvFiles()
}

func chooseCsvFiles() {
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

	for _, choice := range chosenCsvFiles {
		println(choice)
	}
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
