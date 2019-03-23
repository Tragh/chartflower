package main

import (
	"fmt"
	"strconv"
)

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
	var labels []string

	fmt.Println()
	for i, column := range columns {
		fmt.Println(strconv.Itoa(i) + ". " + column + " eg. " + firstRow[i])
	}

	choice := getChoice("Choose label column")

	for i, column := range columns {
		number := strconv.Itoa(i)
		if choice == column {
			labels = getColumnData(i, data)
		} else if choice == number {
			labels = getColumnData(i, data)
		}
	}

	fmt.Println(labels)
	// fmt.Println(firstRow)

	return "bar chart here"
}

func getColumnData(column int, data [][]string) []string {
	var columnData []string
	for i, data := range data {
		if i != 0 {
			columnData = append(columnData, data[column])
		}
	}
	return columnData
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
