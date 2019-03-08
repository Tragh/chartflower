package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	csvFilenames := csvFilesinFolder()
	for _, csvFilename := range csvFilenames {
		println(csvFilename)
	}
}

func csvFilesinFolder() []string {
	var csvFiles []string
	filesInFolder, _ := ioutil.ReadDir("./csv/")
	for _, fileInFolder := range filesInFolder {
		if strings.HasSuffix(fileInFolder.Name(), "csv") {
			csvFiles = append(csvFiles, fileInFolder.Name())
		}
	}
	return csvFiles
}
