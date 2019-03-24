package main

import (
	"fmt"
	"os"
)

func saveFile(chart string) {
	file, _ := os.Create("chart.html")
	defer file.Close()

	file.WriteString(chart)

	fmt.Println("\nSaved chart.html")
}
