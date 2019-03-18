package main

import (
	"fmt"
	"strconv"
	"strings"
)

func chartTypes() [2]string {
	return [2]string{"bar", "line"}
}

func makeChart() {
	selectedChart := chooseChartType()
	for i, chartType := range chartTypes() {
		if selectedChart == strconv.Itoa(i) {
			switch chartType {
			case "bar":
				barChart()
			default:
				fmt.Println("No case for", chartType, "chart")
			}
		}
	}
}

func barChart() {
	data := getData()
	chooseColumns(data)
}

func chooseChartType() string {
	fmt.Println()
	for i, chartType := range chartTypes() {
		fmt.Printf("%v. %v\n", i, chartType)
	}
	fmt.Print("\nChoose chart type:")
	input := getConsoleText()
	choice := strings.Fields(input)
	return choice[0]
}
