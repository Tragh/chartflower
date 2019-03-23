package main

import (
	"fmt"
	"strconv"
	"strings"
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
	columnNames := getColumnNames(data)
	firstRow := getFirstRow(data)
	var labelsIndex int
	var valuesIndex int
	var labels []string
	var values []string

	fmt.Println()
	for i, column := range columnNames {
		fmt.Println(strconv.Itoa(i) + ". " + column + " eg. " + firstRow[i])
	}

	choice := getChoice("Choose label column")

	for i, column := range columnNames {
		number := strconv.Itoa(i)
		if choice == column {
			labelsIndex = i
			labels = getColumnData(i, data)
		} else if choice == number {
			labelsIndex = i
			labels = getColumnData(i, data)
		}
	}

	fmt.Println()
	for i, column := range columnNames {
		if i != labelsIndex {
			fmt.Println(strconv.Itoa(i) + ". " + column + " eg. " + firstRow[i])
		}
	}

	choice = getChoice("Choose values")

	for i, column := range columnNames {
		number := strconv.Itoa(i)
		if choice == column {
			valuesIndex = i
			values = getColumnData(i, data)
		} else if choice == number {
			valuesIndex = i
			values = getColumnData(i, data)
		}
	}

	valuesLabel := columnNames[valuesIndex]
	labelsString := strings.Join(labels, ",")
	valuesString := strings.Join(values, ",")

	return barChartTemplate(labelsString, valuesString, valuesLabel)
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

func getColumnNames(data [][]string) []string {
	var columns []string
	for i, row := range data {
		if i == 0 {
			columns = row
		}
	}
	return columns
}
