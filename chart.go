package main

import "fmt"

func charts() [1]string {
	charts := [1]string{"bar"}
	return charts
}

func makeChart(data [][]string) string {
	var chart string

	switch selectChart() {
	case "bar":
		chart = barChart(data)
	default:
		fmt.Println("No case for", chart)
	}

	return chart
}

func barChart(data [][]string) string {
	columns := getColumns(data)
	firstRow := getFirstRow(data)

	fmt.Println(columns)
	fmt.Println(firstRow)

	return "bar chart here"
}

func getFirstRow(data [][]string) []string {
	var columns []string
	for i, row := range data {
		if i == 1 {
			columns = row
		}
	}
	return columns
}

func getColumns(data [][]string) []string {
	var columns []string
	for i, row := range data {
		if i == 0 {
			columns = row
		}
	}
	return columns
}
