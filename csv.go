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

func selectChart() string {
	var selectedChart string

	fmt.Println()
	for i, chart := range charts() {
		fmt.Println(strconv.Itoa(i) + ". " + chart)
	}

	choice := getChoice("Choose chart")

	for i, chart := range charts() {
		number := strconv.Itoa(i)
		if choice == chart {
			selectedChart = chart
		} else if choice == number {
			selectedChart = chart
		}
	}

	return selectedChart
}

func csvToArray() [][]string {
	var selectedCSV string

	fmt.Println()
	for i, csvFilename := range getCSVFilenames() {
		csvTrimmed := strings.TrimSuffix(csvFilename, ".csv")
		fmt.Println(strconv.Itoa(i) + ". " + csvTrimmed)
	}

	choice := getChoice("Choose csv")

	for i, csvFilename := range getCSVFilenames() {
		csvFilenameWithoutSuffix := strings.TrimSuffix(csvFilename, ".csv")
		number := strconv.Itoa(i)
		if choice == csvFilename {
			selectedCSV = csvFilename
		} else if choice == csvFilenameWithoutSuffix {
			selectedCSV = csvFilename
		} else if choice == number {
			selectedCSV = csvFilename
		}
	}

	data := convertCSVToArray(selectedCSV)
	return data
}

func convertCSVToArray(filename string) [][]string {
	file, err := os.Open("./csv/" + filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
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

func getChoice(prompt string) string {
	fmt.Printf("%v: ", prompt)
	choice := getConsoleText()
	return choice
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
